package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(foodRoutes *gin.Engine) {

	//  Get List of foods
	foodRoutes.GET("/food", controller.GetFoods())
	//  Get Single Food
	foodRoutes.GET("/food/:food_id", controller.GetFood())

	//  Create Food
	foodRoutes.POST("/food", controller.CreateFood())

	//  Update Food usng patch
	foodRoutes.PATCH("/food/:food_id", controller.UpdateFood())

}
