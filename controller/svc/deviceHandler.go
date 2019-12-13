package svc

import (
	"net/http"

	"encoding/json"
	"strconv"

	"github.com/micro/go-micro/util/log"
)

type DeviceInfo struct {
	Guid           int32  `json:"Guid"`
	ProductKey     string `json:"ProductKey"`
	DeviceName     string `json:"DeviceName"`
	DeviceSecret   string `json:"DeviceSecret"`
	Model          string `json:"Model"`
	ProductVersion string `json:"ProductVersion"`
	SdkVersion     string `json:"SdkVersion"`
	Status         int    `json:"Status"`
	CreateAt       string `json:"CreateAt"`
	UpdateAt       string `json:"UpdateAt"`
	LastActiveAt   string `json:"LastActiveAt"`
	ApplyId        string `json:"ApplyId"`
	DeleteFlag     int    `json:"DeleteFlag"`
}

func handleRegisterDevices(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
	productKey := r.URL.Query().Get("ProductKey")
	count, _ := strconv.Atoi(r.Form.Get("Count"))
	// get DeviceNames array from req.body
	log.Infof("register devices product key : %s, count : %d", productKey, count)
	w.Header().Set("Content-Type", "application/json")
	_, err = Ctx.Dbsvc.RegisterDevices(int32(count), productKey)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		type RegDeviceResp struct {
			Code    string
			Message string
			ApplyId string
		}
		var regDeviceResp RegDeviceResp
		regDeviceResp.Code = "200"
		regDeviceResp.ApplyId = ""
		b, _ := json.Marshal(regDeviceResp)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleQueryApplyState(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceListByApplyId(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceInfo(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceList(w http.ResponseWriter, r *http.Request) {
	productKey := r.URL.Query().Get("ProductKey")
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("Offset"))
	keyword := r.URL.Query().Get("Keyword")
	log.Infof("query device list, productKey : %s, limit : %d, offset : %d, keyword : %s", productKey,
		limit, offset, keyword)
	w.Header().Set("Content-Type", "application/json")
	deviceList, err := Ctx.Dbsvc.QueryDeviceList(productKey, int32(offset), int32(limit), keyword)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		type DeviceListResp struct {
			Code           string
			Message        string
			TotalCount     int
			DeviceInfoList []DeviceInfo
		}
		var deviceListResp DeviceListResp
		deviceListResp.Code = "200"
		deviceListResp.TotalCount = len(deviceList)
		deviceListResp.DeviceInfoList = deviceList
		b, _ := json.Marshal(deviceListResp)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleDeleteDevice(w http.ResponseWriter, r *http.Request) {
}

func handleQueryDeviceQuota(w http.ResponseWriter, r *http.Request) {
}

func handleModifyDeviceQuota(w http.ResponseWriter, r *http.Request) {
}

func handleCheckDeviceName(w http.ResponseWriter, r *http.Request) {
}
