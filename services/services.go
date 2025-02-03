package services

import (
	"context"
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ReceiptProcessor struct{}

var receipts = sync.Map{}

func (serv ReceiptProcessor) ProcessReceipt(c context.Context, receipt models.Receipt) string {
	receiptID := uuid.New().String()

	receipts.Store(receiptID, receipt)

	return receiptID
}

func (serv ReceiptProcessor) GetReceiptPoints(c context.Context, receiptID string) (int, bool) {
	var points int
	if value, ok := receipts.Load(receiptID); ok {
		receipt := value.(models.Receipt)
		points = calculatePoints(receipt)
		return points, true

	} else {
		return points, false
	}
}

func calculatePoints(receipt models.Receipt) int {
	points := 0

	reg := regexp.MustCompile("[a-zA-Z0-9]")
	points += len(reg.FindAllString(receipt.Retailer, -1))

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil {
		if total == float64(int(total)) {
			points += 50
		}

		if math.Mod(total, 0.25) == 0 {
			points += 25
		}

		points += (len(receipt.Items) / 2) * 5
	}

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(itemPrice * 0.2))
			}
		}
	}

	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 == 1 {
		points += 6
	}

	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}
