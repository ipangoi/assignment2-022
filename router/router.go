package routers

import (
	"assignment-2/handler"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetOrderRouter() *gin.Engine {
	router := gin.Default()
	controllers := handler.NewOrderHandlerImpl()

	router.GET("/orders/:id", controllers.GetOneOrder)

	router.POST("/orders", controllers.CreateOrder)

	router.GET("/orders", controllers.GetAllOrder)

	router.PATCH("/orders/:id", controllers.UpdateOrder)

	router.DELETE("/orders/:id", controllers.DeleteOrder)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
