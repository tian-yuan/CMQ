package svc

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"github.com/pkg/errors"
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
	logrus.Infof("start mysql server. host : %s, port : %d, database : %s, username : %s, password : %s",
		ds.Conf.Host, ds.Conf.Port, ds.Conf.Database, ds.Conf.Username, ds.Conf.Password)

	addr := net.JoinHostPort(ds.Conf.Host, strconv.Itoa(int(ds.Conf.Port)))
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", ds.Conf.Username, ds.Conf.Password, addr, ds.Conf.Database)
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
	queryStr := fmt.Sprintf("select id, product_key, device_secret from %s where device_name = '%s'", deviceDatabase, deviceName)
	logrus.Infof("query string : %s", queryStr)
	var guid uint32
	var key string
	var deviceSecret string

	rows:= ctx.db.QueryRow(queryStr)
	if rows == nil {
		logrus.Error("query row failed.")
		return 0, errors.New("database internal error.")
	}
	err := rows.Scan(&guid, &key, &deviceSecret)
	if err != nil {
		logrus.Error("query record failed.")
		return 0, err
	}
	logrus.Infof("query record from database, guid : %d, product key : %s", guid, key)
	return guid, nil
}
