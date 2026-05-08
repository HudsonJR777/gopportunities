// @title Goportunities API
// @version 1.0
// @description API for job openings and hiring opportunities
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8081
// @schemes http
// @BasePath /api/v1

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
