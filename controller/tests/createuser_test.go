package tests

import (
	"Golang-MySQL-RestApi-Testing-Docker/controller"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	reqBody, _ := json.Marshal(controller.User{
		Firstname:     "firstname",
		Lastname:      "lastname",
		Referencename: "referencename",
		Country:       "country",
	})
	res, err := http.Post(TEST_APP_URL, "application/json", bytes.NewBuffer(reqBody))

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}
