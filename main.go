package main

import (
	"github.com/Calebbuffleben/api-with-golang/router"
	"github.com/Calebbuffleben/api-with-golang/config"
)

var(
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error? %v", err)
		return
	}

	router.Initialize()
}