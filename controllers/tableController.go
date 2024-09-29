package controllers

import (
	"context"
	"net/http"
	"restaurant_manegment_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()

		tableID := c.Params.ByName("table_id")

		filter := bson.M{"table_id": tableID}
		var table models.Table

		if err := TableCollection.FindOne(ctx, filter).Decode(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Fetching Single Table",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validating Table Struct",
				"Error":   err.Error(),
			})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Table Fetched Successfully",
			"Data":    table,
		})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()
		var table models.Table

		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Binding Request Body To Table Struct",
				"Error":   err.Error(),
			})
		}

		if err := Validate.Struct(table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validating Table Struct ",
				"Error":   err.Error(),
			})
			return

		}

		table.ID = primitive.NewObjectID()
		table.TableID = table.ID.Hex()

		table.CreatedAt = time.Now()
		table.UpdatedAt = time.Now()

		insertTableResult, err := TableCollection.InsertOne(ctx, table)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Insert Table",
				"Error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"Message":     "Table Created Successfully",
			"InsertionID": insertTableResult.InsertedID,
		})

	}
}
