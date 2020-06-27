package libs

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
	"database/sql"
)

var Sqldb *sql.DB
var err error

func Dbintialise() {
	Sqldb, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/refuserinfo")
	if err != nil {
		logger.Error("Error Connecting User Data MYSQL", err)
	} else {
		logger.Info("Connecting DB...")
	}
}
