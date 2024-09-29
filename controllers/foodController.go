package controllers

import (
	"context"
	"math"
	"net/http"
	"restaurant_manegment_api/database"
	"restaurant_manegment_api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var FoodCollection *mongo.Collection = database.OpenCollection(*database.Client, "food")
var MenuCollection *mongo.Collection = database.OpenCollection(*database.Client, "menu")

var Validate *validator.Validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		//  Create Context For Cancellation and Deadline Tracking

		ctx, cancell := context.WithTimeout(context.Background(), 60*time.Second)

		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		foodPerPage := 5
		foodOffset := 0

		if limitStr != "" {
			if limitValue, err := strconv.Atoi(limitStr); err == nil && limitValue >= 1 {
				foodPerPage = limitValue
			}
		}

		if offsetStr != "" {
			if offsetValue, err := strconv.Atoi(limitStr); err == nil && offsetValue >= 0 {
				foodOffset = offsetValue
			}
		}

		defer cancell()

		// Try To Get The Food List Feom The Food Collection
		var foods []models.Food

		filter := bson.M{}
		opts := options.Find().SetLimit(int64(foodPerPage)).SetSkip(int64(foodOffset))

		cursor, err := FoodCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Getting List Of All Foood",
				"Error":   err.Error(),
			})
			return
		}
		if err := cursor.All(ctx, &foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding List Of Food",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "List Of Food Fetched Successfully",
			"Data":    foods,
		})

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()

		foodID := c.Params.ByName("food_id")
		filter := bson.M{
			"food_id": foodID,
		}
		var foundFood models.Food

		if err := FoodCollection.FindOne(ctx, filter).Decode(&foundFood); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding Json Data To Food Struct",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(foundFood); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In ValidateValidate Food",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Food Is Fetched Successfully",
			"Data":    foundFood,
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}

}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food

		// Bind JSON to the Food struct
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error binding JSON to Food struct",
				"Error":   err.Error(),
			})
			return
		}

		// ValidateValidate the Food struct
		if err := Validate.Struct(food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Validation failed for Food struct",
				"Error":   err.Error(),
			})
			return
		}
		var menu models.Menu
		filter := bson.M{"menu_id": food.MenuID}
		if err := MenuCollection.FindOne(ctx, filter).Decode(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Finding Menu With MenuID",
				"Error":   err.Error(),
			})
			return
		}

		// Set timestamps directly
		food.ID = primitive.NewObjectID()
		food.CreatedAt = time.Now()
		var fixedPrice = toFixed(*food.Price, 2)
		food.Price = &fixedPrice
		food.UpdatedAt = time.Now()
		food.FoodID = food.ID.Hex()

		// Insert the food item into the collection
		insertResult, err := FoodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error inserting food",
				"Error":   err.Error(),
			})
			return
		}

		// Return the ID of the created food item as part of the response
		c.JSON(http.StatusCreated, gin.H{
			"Message": "Food created successfully",
			"FoodID":  insertResult.InsertedID,
		})
	}
}
func toFixed(price float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))

	modifiedPrice := math.Round(factor*price) / factor

	return modifiedPrice
}
