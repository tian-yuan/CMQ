package svc

import (
	"k8s.io/apimachinery/pkg/util/json"
	"github.com/sirupsen/logrus"
	"time"
	"fmt"
)

const (
	GUID_PREFIX = "G"
	DEVICE_PREFIX = "D"
)

type DeviceInfo struct{
	DeviceId string     `json:"DI"`
	ProductKey string   `json:"PK"`
}

// G:{guid} = {"DeviceId": "123456", "ProductKey": "key"}
func UpdateDeviceInfo(guid uint32, di DeviceInfo) error {
	key := fmt.Sprintf("%s:%d", GUID_PREFIX, guid)
	dib, err := json.Marshal(di)
	if err != nil {
		logrus.Errorf("update deivcie info for guid : %d, deivce id : %s failed when json format.", guid, di.DeviceId)
		return err
	}
	r := Global.RedisClient.Set(key, string(dib), 30 * 24 * 60 * time.Minute)
	if err = r.Err(); err != nil {
		logrus.Errorf("update deivcie info for guid : %d, deivce id : %s failed for redis update", guid, di.DeviceId)
		return err
	}
	return err
}

func UpdateGuidToDeviceIdMap(guid uint32, dn string) error {
	key := fmt.Sprintf("%s:%s", DEVICE_PREFIX, dn)
	r := Global.RedisClient.Set(key, guid, 30 * 24 * 60 * time.Minute)
	if err := r.Err(); err != nil {
		logrus.Errorf("update guid to device id map failed, guid: %d, deivce id : %s failed for redis update", guid, dn)
		return err
	}
	return nil
}
