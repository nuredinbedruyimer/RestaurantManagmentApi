package controllers

import (
	"context"
	"net/http"
	"restaurant_manegment_api/database"
	"restaurant_manegment_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderItemsCollection *mongo.Collection = database.OpenCollection(*database.Client, "orderItems")

func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()

		orderItemID := c.Params.ByName("order_item_id")
		filter := bson.M{
			"order_item_id": orderItemID,
		}
		var foundOrderItem models.OrderItems

		if err := OrderItemsCollection.FindOne(ctx, filter).Decode(&foundOrderItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding Json Data To OrderItems Struct",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(foundOrderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In ValidateValidate OrderItems",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, foundOrderItem)
	}
}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}

}

func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
