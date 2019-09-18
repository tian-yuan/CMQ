package svc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlSource struct {
	Dsn string
	DB *sql.DB
}

func NewMysqlSource(dsn string) *MysqlSource {
	return &MysqlSource{
		Dsn: dsn,
	}
}

func (ms *MysqlSource) Start() bool {
	db, err := sql.Open("mysql", ms.Dsn)
	if err !=nil {
		return false
	}
	ms.DB = db
	return true
}
