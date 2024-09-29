package controllers

import (
	"context"
	"net/http"
	"restaurant_manegment_api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()
		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		menusPerPage := 5
		menuOffset := 0

		if limitStr != "" {
			if limitValue, err := strconv.Atoi(limitStr); err == nil && limitValue > 1 {
				menusPerPage = limitValue
			}
		}
		if offsetStr != "" {
			if offsetValue, err := strconv.Atoi(offsetStr); err == nil && offsetValue >= 0 {
				menuOffset = offsetValue
			}
		}

		opts := options.Find().SetLimit(int64(menusPerPage)).SetSkip(int64(menuOffset))

		filter := bson.M{}

		cursor, err := MenuCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"Message": "Menus Not Found",
				"Error":   err.Error(),
			})
			return
		}
		var menus []models.Menu

		if err := cursor.All(ctx, &menus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding Fetched Data",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Menus Fetched Successfully",
			"Data":    menus,
		})

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), time.Second*100)

		defer cancell()

		menuID := c.Params.ByName("menu_id")

		filter := bson.M{"menu_id": menuID}
		var menu models.Menu

		if err := MenuCollection.FindOne(ctx, filter).Decode(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Find Menu",
				"Error":   err.Error(),
			})
			return
		}
		if err := Validate.Struct(menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validate Menu Struct",
				"Error":   err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Menu Fetched Successfully",
			"Data":    menu,
		})

	}
}
func IsValidTimeSpan(start time.Time, end time.Time, now time.Time) bool {

	return now.After(start) && now.Before(end)

}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancell()
		var menu models.Menu

		menuID := c.Params.ByName("menu_id")

		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "ErrorIn Binding Request Body to Menu Struct",
				"Error":   err.Error(),
			})
			return
		}
		var updateMonu bson.D
		if menu.StartDate != nil && menu.EndDate != nil {
			if !IsValidTimeSpan(*menu.StartDate, *menu.EndDate, time.Now()) {
				c.JSON(http.StatusBadRequest, gin.H{
					"Message": "The Curret Time is Not Found In The Span",
				})
				return
			}

			updateMonu = append(updateMonu, bson.E{Key: "start_date", Value: menu.StartDate})
			updateMonu = append(updateMonu, bson.E{Key: "end_time", Value: menu.EndDate})
		}
		if menu.Name != nil {
			updateMonu = append(updateMonu, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != nil {
			updateMonu = append(updateMonu, bson.E{Key: "category", Value: menu.Category})

		}
		updateMonu = append(updateMonu, bson.E{Key: "updated_at", Value: menu.UpdatedAt})

		filter := bson.M{"menu_id": menuID}
		upsert := true

		opts := options.UpdateOptions{
			Upsert: &upsert,
		}

		updateMenuResult, err := MenuCollection.UpdateOne(ctx, filter, bson.D{
			bson.E{Key: "$set", Value: updateMonu},
		}, &opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Updating Menu Struct",
				"Error":   err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Message":     "Menu Updated Successfully",
			"UpdatedData": updateMenuResult,
		})

	}

}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancell := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancell()
		var menu models.Menu

		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Binding Request Body To Menu Struct",
				"Error":   err.Error(),
			})
			return
		}
		if err := Validate.Struct(menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": " Error Validation Menu Struct",
				"Error":   err.Error(),
			})
			return
		}
		menu.ID = primitive.NewObjectID()
		menu.CreatedAt = time.Now()
		menu.UpdatedAt = time.Now()
		menu.MenuID = menu.ID.Hex()

		insertMenuResult, err := MenuCollection.InsertOne(ctx, menu)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Inserting Menu",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"Message":     "Menu Created Successfully",
			"InsertionID": insertMenuResult.InsertedID,
		})

	}
}
