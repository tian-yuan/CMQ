package svc

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type DeviceConf struct {
	Host string
	Port uint16
	Password string
	Username string
	Database string
}

type DeviceSvc struct {
	Conf *DeviceConf
}

func NewDeviceConfig() *DeviceConf {
	return &DeviceConf{
		Host:     "",
		Port:     0,
		Password: "",
		Username: "",
		Database: "",
	}
}

func NewDeviceSvc(conf *DeviceConf) *DeviceSvc {
	return &DeviceSvc{
		Conf: conf,
	}
}

func (ds *DeviceSvc) Start() error {
    logrus.WithFields(logrus.Fields{
		"Host": ds.Conf.Host,
		"Port": ds.Conf.Port,
	}).Info("start h2c server.")

	addr := net.JoinHostPort(ds.Conf.Host, strconv.Itoa(int(ds.Conf.Port)))
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@%s/%s", ds.Conf.Username, ds.Conf.Password, addr, ds.Conf.Database)
	db, err := sql.Open("mysql", dsn)
	ctx.db = db
	if err != nil {
		logrus.Error("open mysql connection failed.")
	}
    return err
}

func (ds *DeviceSvc) Stop() {
	ctx.db.Close()
}

const deviceDatabase = "device_info"

func (ds *DeviceSvc) Register(productKey string, deviceName string, sign string) (uint32, error) {
	// query device from database
	queryStr := fmt.Sprintf("select * from %s where device_name = %s", deviceDatabase, deviceName)
	rows, err := ctx.db.Query(queryStr)
	if err != nil {
		logrus.Error("query record failed.")
		return 0, err
	}
	if len(rows) == 0 {
		return 0, nil
	}
}
