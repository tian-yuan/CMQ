package svc

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"strconv"
	"fmt"
)

// iot message dispatcher publish
const publishPath = "/v1/imd/publish"
const kickPath = "/v1/imd/kick"
const healthz = "/v1/imd/healthz"
const rpcPath = "/v1/imd/rpc"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logrus.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == publishPath {
		logrus.Info("publish.")
		handlePublish(w, r)
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
	ctx, errMsg := queryContext(r)
	if ctx != nil {
		deviceName := r.FormValue("deviceName")
		productKey := r.FormValue("productKey")
		msg := r.FormValue("message")
		topic := r.FormValue("topic")
		qosStr := r.FormValue("qos")

		if deviceName == "" || productKey == "" || msg == "" ||
			topic == "" || qosStr == "" {
			logrus.Warningf("Bad publish message request, %v", r)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		qos, e := strconv.Atoi(qosStr)
		if e != nil {
			logrus.Warningf("Bad publish message request, qos is not number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if qos != 0 && qos != 1 {
			logrus.Warningf("Bad publish message request, qos should be 0 or 1")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		code := ctx.PublishMessage(deviceName, productKey, msg, topic, int8(qos))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body := fmt.Sprintf("{\"code\": %d}", code)
		w.Write([]byte(body))
		return
	}

	logrus.Warningf("Bad publish message request, %v, error: %s", r, errMsg)
	w.WriteHeader(http.StatusBadRequest)
}
