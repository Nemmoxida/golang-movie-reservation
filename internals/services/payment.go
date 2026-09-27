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
