package svc

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const controllerPath = "/iothub"
const createProduct = "CreateProduct"
const queryProduct = "QueryProduct"
const queryProductList = "QueryProductList"
const updateProduct = "UpdateProduct"
const deleteProduct = "DeleteProduct"
const queryProductQuota = "QueryProductQuota"
const modifyProductQuota = "ModifyProductQuota"
const checkProductName = "CheckProductName"
const registerDevices = "RegisterDevices"
const queryApplyState = "QueryApplyState"
const queryDeviceListByApplyId = "QueryDeviceListByApplyId"
const queryDeviceInfo = "QueryDeviceInfo"
const queryDeviceList = "QueryDeviceList"
const queryTopicList = "QueryTopicList"
const deleteDevice = "DeleteDevice"
const queryDeviceQuota = "QueryDeviceQuota"
const modifyDeviceQuota = "ModifyDeviceQuota"
const checkDeviceName = "CheckDeviceName"
const createTopicClass = "CreateTopicClass"
const queryTopicClass = "QueryTopicClass"
const queryTopicClassList = "QueryTopicClassList"
const updateTopicClass = "UpdateTopicClass"
const deleteTopicClass = "DeleteTopicClass"
const queryTopicClassQuota = "QueryTopicClassQuota"
const modifyTopicClassQuota = "ModifyTopicClassQuota"
const checkTopicName = "CheckTopicName"
const publishMessage = "PublishMessage"
const queryMessage = "QueryMessage"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logrus.Info(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
	action := r.URL.Query().Get("Action")
	if action == createProduct {
		logrus.Info("create product.")
		handleCreateProduct(w, r)
	} else if action == queryProduct {
		handleQueryProduct(w, r)
	} else if action == queryProductList {
		handleQueryProductList(w, r)
	} else if action == updateProduct {
		handleUpdateProduct(w, r)
	} else if action == deleteProduct {
		handleDeleteProduct(w, r)
	} else if action == queryProductQuota {
		handleQueryProductQuota(w, r)
	} else if action == modifyProductQuota {
		handleModifyProductQuota(w, r)
	} else if action == checkProductName {
		handleCheckProductName(w, r)
	} else if action == registerDevices {
		handleRegisterDevices(w, r)
	} else if action == queryApplyState {
		handleQueryApplyState(w, r)
	} else if action == queryDeviceListByApplyId {
		handleQueryDeviceListByApplyId(w, r)
	} else if action == queryDeviceInfo {
		handleQueryDeviceInfo(w, r)
	} else if action == queryDeviceList {
		handleQueryDeviceList(w, r)
	} else if action == queryTopicList {
		handleQueryTopicList(w, r)
	} else if action == deleteDevice {
		handleDeleteDevice(w, r)
	} else if action == queryDeviceQuota {
		handleQueryDeviceQuota(w, r)
	} else if action == modifyDeviceQuota {
		handleModifyDeviceQuota(w, r)
	} else if action == checkDeviceName {
		handleCheckDeviceName(w, r)
	} else if action == createTopicClass {
		handleCreateTopicClass(w, r)
	} else if action == queryTopicClass {
		handleQueryTopicClass(w, r)
	} else if action == queryTopicClassList {
		handleQueryTopicClassList(w, r)
	} else if action == updateTopicClass {
		handleUpdateTopicClass(w, r)
	} else if action == deleteTopicClass {
		handleDeleteTopicClass(w, r)
	} else if action == queryTopicClassQuota {
		handleQueryTopicClassQuota(w, r)
	} else if action == modifyTopicClassQuota {
		handleModifyTopicClassQuota(w, r)
	} else if action == checkTopicName {
		handleCheckTopicName(w, r)
	} else if action == publishMessage {
		handlePublishMessage(w, r)
	} else if action == queryMessage {
		handleQueryMessage(w, r)
	}
}

func handlePublishMessage(w http.ResponseWriter, r *http.Request) {
}

func handleQueryMessage(w http.ResponseWriter, r *http.Request) {
}
