package base

import (
	"Golang-MySQL-RestApi-Testing-Docker/controller"
	"Golang-MySQL-RestApi-Testing-Docker/libs"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func usercreation() {
	libs.Dbintialise()
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.CreateUsers).Methods("POST")
	router.HandleFunc("/users/{referencename}", controller.UpdateUsers).Methods("PUT")
	router.HandleFunc("/users/{referencename}", controller.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/users/{country}", controller.ListUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8087", router))
}

//user Info
var user chan controller.User

func notifyservice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	var buffer bytes.Buffer
	encode := json.NewEncoder(&buffer)
	encode.Encode(<-user)
	fmt.Fprintf(w, "UserInfo: %v\n\n", buffer.String())
}

func StartUserApplication() {
	usercreation()
}
