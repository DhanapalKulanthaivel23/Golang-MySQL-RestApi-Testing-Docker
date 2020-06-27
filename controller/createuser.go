package controller

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs"
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	smt := libs.Sqldb
	stmt, err := smt.Prepare("INSERT INTO users(firstname,lastname,referencename,country) VALUES(?,?,?,?)")
	if err != nil {
		logger.Error("DB Error while Insert User", err)
	} else {
		logger.Info("New USer Created")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error on user data", err)
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	referencename := keyVal["referencename"]
	country := keyVal["country"]
	_, err = stmt.Exec(firstname, lastname, referencename, country)
	if err != nil {
		logger.Error("Error on User Data Execution", err)
	}
	fmt.Fprintf(w, "User %s Created", referencename)
}
