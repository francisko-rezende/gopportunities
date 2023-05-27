package main

import (
	"github.com/francisko-rezende/gopportunities/config"
	"github.com/francisko-rezende/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize db
	err := config.Init()

	if err != nil {
		logger.ErrF("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Init()
}
