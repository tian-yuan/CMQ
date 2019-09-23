package svc

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const registerPath = "/v1/device/register"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logrus.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	u := r.URL.EscapedPath()
	if u == registerPath{
		logrus.Info("subscribe.")
		handleRegister(w, r)
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	deviceName := r.Form.Get("DeviceName")
	productKey := r.Form.Get("ProductKey")
	logrus.Infof("handle register deviceName : %s, productKey: %s", deviceName, productKey)
}
