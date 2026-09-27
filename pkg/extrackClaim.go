package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// unused for now

func ExtrackClaims(c *gin.Context) *Claims {
	claimsValue, exist := c.Get("claims")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})

	}

	claims := claimsValue.(*Claims)
	return claims
}
