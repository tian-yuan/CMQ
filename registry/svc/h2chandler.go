package svc

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
)

const registerPath = "/v1/device/register"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == registerPath {
		log.Info("subscribe.")
		handleRegister(w, r)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	deviceName := r.Form.Get("DeviceName")
	productKey := r.Form.Get("ProductKey")
	log.Infof("handle register deviceName : %s, productKey: %s", deviceName, productKey)
}
