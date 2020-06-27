package controller

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs"
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := libs.Sqldb.Prepare("DELETE FROM users WHERE referencename = ?")
	if err != nil {
		logger.Error("DB Error Removing User Data", err)
	} else {
		logger.Info("USer Removed...")
	}
	_, err = stmt.Exec(params["referencename"])
	if err != nil {
		logger.Error("Error Removing User Data Execution", err)
	}
	fmt.Fprintf(w, "User %s deleted", params["referencename"])
}
