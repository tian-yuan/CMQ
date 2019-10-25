package svc

import (
	"github.com/sirupsen/logrus"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var ctx Context

func init() {
}

type Context struct {
	db *sql.DB
}

func (ctx *Context) register() error {
	logrus.Info("register.")
	return nil
}
