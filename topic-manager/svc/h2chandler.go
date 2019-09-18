package svc

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const subscribePath = "/v1/topic/subscribe"
const publishPath = "/v1/topic/publish"

// load topic from database to cache
const topicLoad = "/v1/topic/load"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logrus.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == subscribePath {
		logrus.Info("subscribe.")
		handleSubscribe(w, r)
	} else if u == publishPath {
		handlePublish(w, r)
	} else if u == topicLoad {
		handleTopicLoad(w, r)
	}
}

func handleSubscribe(w http.ResponseWriter, r *http.Request) {
	topic := r.Form.Get("Topic")
	qos := r.Form.Get("Qos")
	guid := r.Form.Get("Guid")
	logrus.Infof("handle subscribe topic : %s, qos : %d", topic, qos)
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
	topic := r.Form.Get("Topic")
	qos := r.Form.Get("Qos")
	logrus.Infof("handle publish topic : %s, qos : %d", topic, qos)
	subs := ctx.match(topic)
	for _, sub := range subs {
		w.Write()
	}
}

func handleTopicLoad(w http.ResponseWriter, r *http.Request) {
	// load subscribe topic
}
