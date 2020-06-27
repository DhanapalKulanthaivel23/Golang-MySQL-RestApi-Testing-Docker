package controller

type User struct {
	Firstname     string `json : "firstname"`
	Lastname      string `json : "lastname"`
	Referencename string `json : "referencename"`
	Country       string `json : "country"`
}
