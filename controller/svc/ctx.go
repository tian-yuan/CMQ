package svc

import (
	_ "github.com/go-sql-driver/mysql"
)

var Ctx Context

func init() {
}

type Context struct {
	Dbsvc *DatabaseSvc
}

type ErrInfo struct {
	Code string `json:"Code"`
	Message string `json:"Message"`
}