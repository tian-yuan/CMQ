package svc

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/google/uuid"
	"encoding/json"
	"bytes"
)

type ProductInfo struct {
	ProductKey string `json:"ProductKey"`
	ProductName string `json:"ProductName"`
	Description string `json:"Description"`
	DeviceCount int32 `json:"DeviceCount"`
	AccessPoints string `json:"AccessPoints"`
	CreateAt string `json:"CreateAt"`
	UpdateAt string `json:"UpdateAt"`
}

func handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.Form.Get("ProductName")
	productDesc := r.Form.Get("Description")
	// create product
	productInfo := &ProductInfo{
		ProductKey: uuid.New().String(),
		ProductName: productName,
		Description: productDesc,
		DeviceCount: 0,
		AccessPoints: "",
		CreateAt: "",
		UpdateAt: "",
	}
	logrus.Infof("create product name : %s, description : %s", productInfo.ProductName, productInfo.Description)
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(productInfo)
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func handleQueryProduct(w http.ResponseWriter, r *http.Request) {
}

func handleQueryProductList(w http.ResponseWriter, r *http.Request) {
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
}

func handleQueryProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleModifyProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleCheckProductName(w http.ResponseWriter, r *http.Request) {
}
