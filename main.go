package main

import (
	"receipt-processor/controllers"

	"github.com/gin-gonic/gin"
)

var (
	ctrl = new(controllers.ReceiptProcessor)
)

func main() {
	r := gin.Default()
	port := "8080"

	r.POST("/receipts/process", ctrl.ProcessReceipt)
	r.GET("/receipts/:id/points", ctrl.GetReceiptPoints)
	r.Run(":" + port)
}
