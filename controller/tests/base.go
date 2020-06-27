package tests

import (
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils"
)

var TEST_APP_URL = utils.GetEnvVar("TEST_APP_URL", "http://localhost:8087/users")
