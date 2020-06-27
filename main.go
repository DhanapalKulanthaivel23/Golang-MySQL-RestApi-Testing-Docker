package main

import (
	"Golang-MySQL-RestApi-Testing-Docker/base"
	"Golang-MySQL-RestApi-Testing-Docker/libs/utils/logger"
)

func main() {
	logger.Info("Starting User Application ...")
	base.StartUserApplication()
}
