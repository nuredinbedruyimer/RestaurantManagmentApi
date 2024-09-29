package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(noteRoute *gin.Engine) {
	//  Get List Of OrderItems
	// noteRoute.GET("/orderItems", controller.GetOrderItems())
	//  Get Single OrderItems
	noteRoute.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	//  Get Order Items Usig Order
	// noteRoute.GET("/orderItems/:order_id", controller.GetOrderItemsByOrder())

	//  Update OrderItems using Patch
	// noteRoute.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

	//  Create OrderItems
	// noteRoute.POST("/orderItems", controller.CreateOrderItem())

}
