package routes

import "github.com/gin-gonic/gin"

func MenuRoutes(menuRoute *gin.Engine) {
	//  Get List Of Menus
	menuRoute.GET("/menus", controller.GetMenus())
	//  Get Single Menu
	menuRoute.GET("/menus/:menu_id", controller.GetMenu())
	//  Create Menu
	menuRoute.POST("/menus", controller.CreateMenu())
	//  update Menu
	menuRoute.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
