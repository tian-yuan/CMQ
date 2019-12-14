package svc

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/RoaringBitmap/roaring"

	"github.com/micro/go-micro/util/log"
)

const subscribePath = "/v1/topic/subscribe"
const publishPath = "/v1/topic/publish"

// load topic from database to cache
const topicLoad = "/v1/topic/load"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof("Hello, %s, http: %d", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == subscribePath {
		log.Info("subscribe.")
		handleSubscribe(w, r)
	} else if u == publishPath {
		handlePublish(w, r)
	} else if u == topicLoad {
		handleTopicLoad(w, r)
	} else {
		log.Errorf("unkown path : %s", u)
	}
}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	topic := r.Form.Get("Topic")
	qos, _ := strconv.Atoi(r.Form.Get("Qos"))
	guid, _ := strconv.ParseUint(r.Form.Get("Guid"), 10, 64)
	log.Infof("handle subscribe topic : %s, qos : %d, guid : %d", topic, qos, guid)
	err := Ctx.Subscribe(topic, qos, uint32(guid))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
	topic := r.Form.Get("Topic")
	qos := r.Form.Get("Qos")
	log.Infof("handle publish topic : %s, qos : %d", topic, qos)
	subs := Ctx.Match(topic)
	rb := roaring.BitmapOf()
	for _, sub := range subs {
		rb.Add(sub.(uint32))
	}
	buf := new(bytes.Buffer)
	rb.WriteTo(buf)
	w.Write(buf.Bytes())
	w.WriteHeader(http.StatusOK)
}

func handleTopicLoad(w http.ResponseWriter, r *http.Request) {
	// load subscribe topic
}
