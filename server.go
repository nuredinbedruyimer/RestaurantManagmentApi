package main

import (
	"restaurant_manegment_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.NoteRoutes(router)
	routes.TableRoutes(router)

	router.Run(":8888")

}
