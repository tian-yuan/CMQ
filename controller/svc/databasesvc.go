package svc

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"

	"errors"

	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
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
	addr := net.JoinHostPort(ds.Conf.Host, strconv.Itoa(int(ds.Conf.Port)))
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", ds.Conf.Username, ds.Conf.Password, addr, ds.Conf.Database)
	db, err := sql.Open("mysql", dsn)
	ds.Db = db
	if err != nil {
		log.Error("open mysql connection failed.")
	}
	log.Info("open mysql connection success.")
	return err
}

func (ds *DatabaseSvc) Stop() {
	ds.Db.Close()
}

const productDatabase = "product_info"

func (ds *DatabaseSvc) CreateProduct(info ProductInfo) (uint32, error) {
	// create product
	createStr := fmt.Sprintf("insert into %s(product_key, product_name, product_secret, description, "+
		"access_points, device_count, delete_flag) "+
		"values(?, ?, ?, ?, ?, ?, ?)", productDatabase)
	stmtIns, err := ds.Db.Prepare(createStr)
	if err != nil {
		log.Error("database prepare failed.")
		return 0, err
	}
	defer stmtIns.Close()
	var res sql.Result
	res, err = stmtIns.Exec(info.ProductKey, info.ProductName, info.ProductSecret, info.Description, info.AccessPoints,
		info.DeviceCount, info.DeleteFlag)
	if err != nil {
		log.Error("create product failed.")
		return 0, err
	}
	guid, err := res.LastInsertId()
	if err != nil {
		log.Error("fetch last insert id failed.")
		return 0, err
	}
	return uint32(guid), nil
}

func (ds *DatabaseSvc) QueryProductInfo(productKey string) (*ProductInfo, error) {
	// query product
	queryStr := fmt.Sprintf("select product_name, description, access_points, device_count, create_at, update_at "+
		"from %s where product_key='%s'", productDatabase, productKey)
	log.Infof("query string : %s", queryStr)

	rows := ds.Db.QueryRow(queryStr)
	if rows == nil {
		log.Error("query row failed.")
		return nil, errors.New("database internal error.")
	}
	var productInfo ProductInfo
	err := rows.Scan(&productInfo.ProductName, &productInfo.Description, &productInfo.AccessPoints,
		&productInfo.DeviceCount, &productInfo.CreateAt, &productInfo.UpdateAt)
	if err != nil {
		log.Error("query record failed.")
		return nil, err
	}
	log.Infof("query record from database, product key : %s", productKey)
	return &productInfo, nil
}

func (ds *DatabaseSvc) QueryProductList(offset int32, limit int32, keyword string) ([]ProductInfo, error) {
	var productInfoList []ProductInfo
	queryStr := fmt.Sprintf("select product_key, product_name, description, access_points, device_count, create_at, update_at "+
		"from %s order by id limit %d offset %d", productDatabase, limit, offset)
	rows, err := ds.Db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var productInfo ProductInfo
		err := rows.Scan(&productInfo.ProductKey, &productInfo.ProductName, &productInfo.Description, &productInfo.AccessPoints,
			&productInfo.DeviceCount, &productInfo.CreateAt, &productInfo.UpdateAt)
		if err != nil {
			return nil, err
		}
		productInfoList = append(productInfoList, productInfo)
	}

	return productInfoList, nil
}

////////////////////////////////////////////////
const deviceDatabase = "device_info"

func (ds *DatabaseSvc) RegisterDevices(count int32, productKey string) (string, error) {
	var i int32
	for i = 0; i < count; i++ {
		deviceName := uuid.New().String()
		deviceSecret := uuid.New().String()
		createStr := fmt.Sprintf("insert into %s(product_key, device_name, device_secret) values(?, ?, ?)",
			deviceDatabase)
		stmtIns, err := ds.Db.Prepare(createStr)
		if err != nil {
			log.Error("database prepare failed.")
			return "", err
		}
		defer stmtIns.Close()
		_, err = stmtIns.Exec(productKey, deviceName, deviceSecret)
		if err != nil {
			log.Error("create product failed.")
			return "", err
		}
		return "", nil
	}
	return "", nil
}

func (ds *DatabaseSvc) QueryDeviceList(productKey string, offset int32, limit int32, keyword string) ([]DeviceInfo, error) {
	var deviceInfoList []DeviceInfo
	queryStr := fmt.Sprintf("select product_key, device_name, device_secret, model, product_version, sdk_version, create_at, update_at, "+
		"ifnull(last_active_at, ''), ifnull(apply_id, ''), status, delete_flag from %s where product_key='%s' order by id limit %d offset %d",
		deviceDatabase, productKey, limit, offset)
	rows, err := ds.Db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var deviceInfo DeviceInfo
		err := rows.Scan(&deviceInfo.ProductKey, &deviceInfo.DeviceName, &deviceInfo.DeviceSecret, &deviceInfo.Model,
			&deviceInfo.ProductVersion, &deviceInfo.SdkVersion, &deviceInfo.CreateAt, &deviceInfo.UpdateAt,
			&deviceInfo.LastActiveAt, &deviceInfo.ApplyId, &deviceInfo.Status, &deviceInfo.DeleteFlag)
		if err != nil {
			return nil, err
		}
		deviceInfoList = append(deviceInfoList, deviceInfo)
	}

	return deviceInfoList, nil
}
