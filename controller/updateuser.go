package controller

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs"
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := libs.Sqldb.Prepare("UPDATE users SET firstname=?,lastname=?, country=? WHERE referencename = ?")
	if err != nil {
		logger.Error("DB Error on Update User Data", err)
	} else {
		logger.Info("USer Updated...")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error Read User Data", err)
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	country := keyVal["country"]
	_, err = stmt.Exec(firstname, lastname, country, params["referencename"])
	if err != nil {
		logger.Error("Error Update User Data Execution", err)
	}
	fmt.Fprintf(w, "User %s Updated", params["referencename"])
}
