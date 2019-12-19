package svc

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// iot message dispatcher publish
const publishPath = "/v1/imd/publish"
const kickPath = "/v1/imd/kick"
const healthz = "/v1/imd/healthz"
const rpcPath = "/v1/imd/rpc"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof("Hello, %s, http: %d", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == publishPath {
		log.Info("publish.")
		handlePublish(w, r)
	} else {
		log.Errorf("unknown path : %s", u)
	}
}

func queryContext(r *http.Request) (*ClientCtx, string) {
	fd := r.FormValue("fd")
	if fd == "" {
		return nil, "Fd is empty"
	}

	fdi, e := strconv.Atoi(fd)
	if e == nil {
		ctx := GetCtxForFd(int64(fdi))
		if ctx != nil {
			return ctx, ""
		} else {
			return nil, "No fd found"
		}
	} else {
		return nil, "Fd is not number"
	}
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header),
	)
	if err != nil {
		// Optionally record something about err here
	}

	// Create the span referring to the RPC client if available.
	// If wireContext == nil, a root span will be created.
	serverSpan := opentracing.StartSpan(
		"IotHub-HandlePublish",
		ext.RPCServerOption(wireContext),
	)

	defer serverSpan.Finish()

	ctx, errMsg := queryContext(r)
	if ctx != nil {
		deviceName := r.FormValue("deviceName")
		productKey := r.FormValue("productKey")
		msg := r.FormValue("message")
		topic := r.FormValue("topic")
		qosStr := r.FormValue("qos")

		if msg == "" ||
			topic == "" || qosStr == "" {
			log.Infof("Bad publish message request, %v", r)
			serverSpan.SetTag("message", "bad request.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		qos, e := strconv.Atoi(qosStr)
		if e != nil {
			log.Infof("Bad publish message request, qos is not number")
			serverSpan.SetTag("message", "bad request.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if qos != 0 && qos != 1 {
			log.Infof("Bad publish message request, qos should be 0 or 1")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		code := ctx.PublishMessage(deviceName, productKey, msg, topic, int8(qos))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body := fmt.Sprintf("{\"code\": %d}", code)
		serverSpan.SetTag("message", "publish message to device success.")
		w.Write([]byte(body))
		return
	} else {
		serverSpan.SetTag("message", "find device connection info failed.")
	}

	log.Infof("Bad publish message request, %v, error: %s", r, errMsg)
	w.WriteHeader(http.StatusBadRequest)
}
