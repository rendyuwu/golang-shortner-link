package app

import (
	"database/sql"
	"time"

	"github.com/rendyuwu/golang-shortner-link/env"
	"github.com/rendyuwu/golang-shortner-link/helper"
)

func NewDB(env env.MyEnv) *sql.DB {
	db, err := sql.Open("mysql", ""+env["Username"]+":"+env["Password"]+"@tcp("+env["Host"]+":"+env["Port"]+")/"+env["Database"]+"")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(150)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
