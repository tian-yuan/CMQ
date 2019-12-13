package svc

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/util/log"
)

var ctx Context

func init() {
}

type Context struct {
	db *sql.DB
}

func (ctx *Context) register() error {
	log.Info("register.")
	return nil
}
