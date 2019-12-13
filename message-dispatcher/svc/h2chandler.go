package svc

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
)

// iot message dispatcher publish
const publishPath = "/v1/imd/publish"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == publishPath {
		log.Info("publish.")
		handlePublish(w, r)
	}
}

func handlePublish(w http.ResponseWriter, r *http.Request) {
}
