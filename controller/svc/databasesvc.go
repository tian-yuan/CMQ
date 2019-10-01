package svc

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
)

type DatabaseConf struct {
	Host     string
	Port     uint16
	Password string
	Username string
	Database string
}

type DatabaseSvc struct {
	Conf *DatabaseConf
	Db   *sql.DB
}

func NewDatabaseConfig() *DatabaseConf {
	return &DatabaseConf{
		Host:     "",
		Port:     0,
		Password: "",
		Username: "",
		Database: "",
	}
}

func NewDatabaseSvc(conf *DatabaseConf) *DatabaseSvc {
	return &DatabaseSvc{
		Conf: conf,
	}
}

func (ds *DatabaseSvc) Start() error {
	logrus.WithFields(logrus.Fields{
		"Host": ds.Conf.Host,
		"Port": ds.Conf.Port,
	}).Info("start h2c server.")

	addr := net.JoinHostPort(ds.Conf.Host, strconv.Itoa(int(ds.Conf.Port)))
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@%s/%s", ds.Conf.Username, ds.Conf.Password, addr, ds.Conf.Database)
	db, err := sql.Open("mysql", dsn)
	ds.Db = db
	if err != nil {
		logrus.Error("open mysql connection failed.")
	}
	return err
}

func (ds *DatabaseSvc) Stop() {
	ds.Db.Close()
}

const deviceDatabase = "device_info"

func (ds *DatabaseSvc) CreateProduct(info ProductInfo) (uint32, error) {
	// create product
	createStr := fmt.Sprintf("insert into %s(product_key, product_name, product_secret, description, "+
		"access_points, device_count, create_at, update_at, delete_flag) "+
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?)", deviceDatabase)
	stmtIns, err := ds.Db.Prepare(createStr)
	if err != nil {
		logrus.Error("database prepare failed.")
		return 0, err
	}
	defer stmtIns.Close()
	var res sql.Result
	res, err = stmtIns.Exec(info.ProductKey, info.ProductName, info.ProductSecret, info.Description, info.AccessPoints,
		info.DeviceCount, info.CreateAt, info.UpdateAt, info.DeleteFlag)
	if err != nil {
		logrus.Error("create product failed.")
		return 0, err
	}
	guid, err := res.LastInsertId()
	if err != nil {
		logrus.Error("fetch last insert id failed.")
		return 0, err
	}
	return uint32(guid), nil
}
