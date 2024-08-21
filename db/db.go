package db

import (
	"database/sql"
	"rest-api-echo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	conncectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_DATABASE

	db, err = sql.Open("mysql", conncectionString)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Valid")
	}
}

func CreateCon() *sql.DB {
	return db
}
