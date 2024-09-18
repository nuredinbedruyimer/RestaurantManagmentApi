package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(orderRoute *gin.Engine) {
	//  Get All orders
	orderRoute.GET("/orders", controller.GetOrders())
	//  Get Single order
	orderRoute.GET("/orders/:order_id", controller.GetOrder())

	//  Create Order Using Post
	orderRoute.POST("/orders", controller.CreateOrder())

	//  Update Orders Using Patch
	orderRoute.PATCH("/orders/order_id", controller.UpdateOrder())

}
