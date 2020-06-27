package controller

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs"
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var users []User
	result, err := libs.Sqldb.Query("SELECT firstname, lastname, referencename, country FROM users WHERE country = ?", params["country"])
	if err != nil {
		logger.Error("DB Error on Read User Data", err)
	} else {
		logger.Info("USers Listed...")
	}
	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.Firstname, &user.Lastname, &user.Referencename, &user.Country)
		if err != nil {
			logger.Error("Error on result User Data", err)
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
