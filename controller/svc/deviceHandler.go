package svc

import (
	"net/http"
	"github.com/sirupsen/logrus"
)

type DeviceInfo struct {
	ProductKey string `json:"ProductKey"`
	DeviceName string `json:"DeviceName"`
	DeviceSecret string `json:"DeviceSecret"`
	State string `json:"State"`
	CreateAt string `json:"CreateAt"`
	UpdateAt string `json:"UpdateAt"`
	LastActiveAt string `json:"LastActiveAt"`
}

func handleRegisterDevices(w http.ResponseWriter, r *http.Request) {
	productKey := r.Form.Get("ProductKey")
	count := r.Form.Get("Count")
	// get DeviceNames array from req.body
	logrus.Infof("register devices product key : %s, count : %s", productKey, count)
}

func handleQueryApplyState(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceListByApplyId(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceInfo(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceList(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteDevice(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceQuota(w http.ResponseWriter, r *http.Request) {
}

func handleModifyDeviceQuota(w http.ResponseWriter, r *http.Request) {
}

func handleCheckDeviceName(w http.ResponseWriter, r *http.Request) {
}
