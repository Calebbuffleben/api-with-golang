package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.Use(gin.Logger())
	initializeRoutes(router)

	router.Run(":8080") // This is the correct syntax to run the server on port 8080.
}