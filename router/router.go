package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":8081") // listens on 0.0.0.0:8081 by default
}
