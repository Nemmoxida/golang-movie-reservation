package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MakePayment(c *gin.Context) {
	// snap initialization
	var s snap.Client
	s.New(os.Getenv("MIDTRANS_CLIENT"), midtrans.Sandbox)

	uuid := uuid.New().String()

	// create snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid,
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": snapResp.Token, "redirect": snapResp.RedirectURL})
}

// func MidtransWebhook(c *gin.Context) {
// 	var notification struct {
// 		OrderID           string `json:"order_id"`
// 		TransactionID     string `json:"transaction_id"`
// 		TransactionStatus string `json:"transaction_status"`
// 		StatusCode        string `json:"status_code"`
// 		GrossAmount       string `json:"gross_amount"`
// 		SignatureKey      string `json:"signature_key"`
// 		FraudStatus       string `json:"fraud_status"`
// 	}

// 	if err := c.ShouldBindJSON(&notification); err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// verify signature
// 	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

// 	input := notification.OrderID +
// 		notification.StatusCode +
// 		notification.GrossAmount +
// 		serverKey

// 	hash := sha512.Sum512([]byte(input))

// 	expectedSignature := hex.EncodeToString(hash[:])

// 	if expectedSignature != notification.SignatureKey {
// 		c.JSON(401, gin.H{
// 			"error": "invalid signature",
// 		})
// 		return
// 	}

// 	if notification.TransactionStatus == "settlement" {
// 		// update database
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "notification received",
// 	})
// }
