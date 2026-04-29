package main

import (
	"github.com/HudsonJR777/gopportunities/config"
	"github.com/HudsonJR777/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	error := config.Init()

	if error != nil {
		logger.Errorf("Config initialization error: %v", error)
		return
	}

	router.Initialize()
}
