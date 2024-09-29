package controllers

import (
	"context"
	"net/http"
	"restaurant_manegment_api/database"
	"restaurant_manegment_api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var OrderCollection *mongo.Collection = database.OpenCollection(*database.Client, "order")
var TableCollection *mongo.Collection = database.OpenCollection(*database.Client, "table")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)
		defer cancell()

		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		orderPerPage := 5
		orderOffset := 0

		if limitStr != "" {
			if limitValue, err := strconv.Atoi(limitStr); err == nil && limitValue >= 1 {
				orderPerPage = limitValue
			}

		}
		if offsetStr != "" {
			if offsetValue, err := strconv.Atoi(offsetStr); err == nil && offsetValue >= 0 {
				orderOffset = offsetValue
			}

		}

		var orders []models.Order
		filter := bson.M{}
		opts := options.Find().SetLimit(int64(orderPerPage)).SetSkip(int64(orderOffset))

		cursor, err := OrderCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Fetch List Of Orders",
				"Error ":  err.Error(),
			})
			return
		}
		if err := cursor.All(ctx, &orders); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding The Orders Cursor",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Orders Fetched Successfully !!",
			"Data":    orders,
		})

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)
		defer cancell()

		var order models.Order

		orderID := c.Params.ByName("order_id")

		filter := bson.M{"order_id": orderID}
		if err := OrderCollection.FindOne(ctx, filter).Decode(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding Order",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validating Order",
				"Error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "Order Fetched Successfully",
			"Data":    order,
		})

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()
		var order models.Order

		//  Bind The Request Body To Order Struct

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Binding Request Body",
				"Error":   err.Error(),
			})

			return
		}

		if err := Validate.Struct(order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validating order Struct",
				"Error":   err.Error(),
			})
			return
		}
		var table models.Table

		filter := bson.M{"table_id": order.TableID}

		if err := TableCollection.FindOne(ctx, filter).Decode(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Fetching and Decoding The table struct",
				"Error":   err.Error(),
			})
			return
		}

		order.ID = primitive.NewObjectID()
		order.OrderID = order.ID.Hex()
		order.CreatedAt = time.Now()
		order.UpdatedAt = time.Now()

		insertOrderResult, err := OrderCollection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Inserting Order",
				"Error":   err.Error(),
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"Message":     "Order Created Successfully",
			"InsertionID": insertOrderResult.InsertedID,
		})

	}
}
