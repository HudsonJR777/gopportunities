package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
