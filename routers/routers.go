package routers

import (
	"assignment2/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/orders/:id", controllers.GetOrderByID)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	return router
}
