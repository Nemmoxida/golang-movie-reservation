package middleware

import (
	"golang-movie-reservation/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandlerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unautorized": "No Token Found"})
			return
		}

		claims, err := pkg.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}

		if claims.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unauthorize", "message": "admin only route"})
			return
		}

		c.Set("claims", claims)
		c.Next()

	}

}
func AuthHandlerClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unautorized": "No Token Found"})
			return
		}

		claims, err := pkg.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}

		if claims.Role != "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unauthorize", "message": "no role detected"})
			return
		}

		c.Set("claims", claims)
		c.Next()

	}

}
