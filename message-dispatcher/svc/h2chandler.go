package svc

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// iot message dispatcher publish
const publishPath = "/v1/imd/publish"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logrus.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == publishPath {
		logrus.Info("publish.")
		handlePublish(w, r)
	}
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
}
