package main

import (
	"os"
	"restaurant_manegment_api/database"
	"restaurant_manegment_api/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollections *mongo.Collection = database.OpenCollection(*database.Client, "food")

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.NoteRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + PORT)

}
