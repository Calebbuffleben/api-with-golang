package router

import (
	"github.com/Calebbuffleben/api-with-golang/handler"             
	"github.com/gin-gonic/gin"
	"github.com/Calebbuffleben/api-with-golang/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET( "/opening", handler.ShowOpeningHandler)
		v1.POST( "/opening", handler.CreateOpeningHandler)
		v1.DELETE( "/opening", handler.DeleteOpeningHandler)
		v1.PUT( "/opening", handler.UpdateOpeningHandler)
		v1.GET( "/openings", handler.ListOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}