package routes

import (
	controller "restaurant_manegment_api/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(invoicesRoute *gin.Engine) {
	//  Get List Of Invoices
	invoicesRoute.GET("/invoices", controller.GetInvoices())
	// Get Single Invoice
	invoicesRoute.GET("/invoices/:invoice_id", controller.GetInvoice())
	//  Create Invoices
	invoicesRoute.POST("/invoices", controller.CreateInvoice())
	//  Update Invoices using Patch
	invoicesRoute.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())

}
