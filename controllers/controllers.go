package controllers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
)

type ReceiptProcessor struct{}

var (
	serv = new(services.ReceiptProcessor)
)

func (ctrl ReceiptProcessor) ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	receiptID := serv.ProcessReceipt(c, receipt)

	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func (ctrl ReceiptProcessor) GetReceiptPoints(c *gin.Context) {
	receiptID := c.Param("id")

	points, receiptFound := serv.GetReceiptPoints(c, receiptID)

	if receiptFound {
		c.JSON(http.StatusOK, gin.H{"points": points})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
	}
}
