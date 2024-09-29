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

var InvoiceCollection *mongo.Collection = database.OpenCollection(*database.Client, "incoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancell()

		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		invoicePerPage := 5
		invoiceOffset := 0

		if limitStr != "" {
			if limitValue, err := strconv.Atoi(limitStr); err == nil && limitValue >= 1 {
				invoicePerPage = limitValue
			}
		}

		if offsetStr != "" {
			if offsetValue, err := strconv.Atoi(offsetStr); err == nil && offsetValue >= 0 {
				invoiceOffset = offsetValue
			}

		}

		opts := options.Find().SetLimit(int64(invoicePerPage)).SetSkip(int64(invoiceOffset))

		filter := bson.M{}

		cursor, err := InvoiceCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Fetching List Of Invoices",
				"Error":   err.Error(),
			})
			return
		}

		var invoices []models.Invoice

		if err := cursor.All(ctx, &invoices); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Decoding Invoices",
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "List Of Invoices Fetched Successfully",
			"Data":    invoices,
		})

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancell()

		invoiceID := c.Params.ByName("invoice_id")
		filter := bson.M{"invoice_id": invoiceID}
		var invoice models.Invoice

		if err := InvoiceCollection.FindOne(ctx, filter).Decode(&invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Fetching Invoice",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In validating Single Inoice",
				"Error":   err.Error(),
			})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"Message": "Invoice Is Fetche Successfully",
			"Data":    invoice,
		})

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancell := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancell()
		var invoice models.Invoice

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Binding Request Body To Invoice Struct",
				"Error":   err.Error(),
			})
			return
		}

		if err := Validate.Struct(invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Validating Invoice Struct",
				"Error":   err.Error(),
			})
			return
		}
		var order models.Order

		filter := bson.M{"order_id": invoice.OrderID}

		if err := OrderCollection.FindOne(ctx, filter).Decode(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error In Finding Order ",
				"Error":   err.Error(),
			})
			return
		}

		invoice.ID = primitive.NewObjectID()
		invoice.InvoiceID = invoice.ID.Hex()
		invoice.CreatedAt = time.Now()

		invoice.UpdatedAt = time.Now()
		invoice.PaymentDueDate = time.Now()

		insertInvoiceResult, err := InvoiceCollection.InsertOne(ctx, invoice)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Error In Inserting Invoice",
				"Error":   err.Error(),
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"Message":     "Invoice Created Successfully",
			"InsertionID": insertInvoiceResult.InsertedID,
		})

	}
}
