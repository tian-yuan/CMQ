package svc

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
)

// iot business manager publish
const publishPath = "/v1/imd/publish"
const kickPath = "/v1/imd/kick"
const healthz = "/v1/imd/healthz"
const rpcPath = "/v1/imd/rpc"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof("Hello, %s, http: %d", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == publishPath {
		log.Info("publish.")
	} else {
		log.Errorf("unknown path : %s", u)
	}
}
