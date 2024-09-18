package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(tableRoute *gin.Engine) {
	//  Get All List Of Tables Using GetMethod
	tableRoute.GET("/tables", controller.GetTables())
	//  Get Single Table Using Get With Param
	tableRoute.GET("/tables/:table_id", controller.GetTable())
	//  Create Table Using Post Method
	tableRoute.POST("/tables", controller.CreateTable())
	//  Update Table Using Patch Method
	tableRoute.PATCH("/tables/:table_id", controller.UpdateTable())

}
