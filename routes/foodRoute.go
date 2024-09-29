package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(foodRoutes *gin.Engine) {

	//  Get List of foods
	foodRoutes.GET("/foods", controller.GetFoods())
	//  Get Single Food
	foodRoutes.GET("/foods/:food_id", controller.GetFood())

	//  Create Food
	foodRoutes.POST("/foods", controller.CreateFood())

	//  Update Food usng patch
	foodRoutes.PATCH("/foods/:food_id", controller.UpdateFood())

}
