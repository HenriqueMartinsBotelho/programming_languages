package main

import (
	"aprendendo_go/project_1/config"
	"aprendendo_go/project_1/router"
)

var (
	logger config.Logger
)


func main() {
	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}