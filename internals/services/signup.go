package services

import (
	"context"
	"golang-movie-reservation/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var req UserReq

	c.ShouldBindJSON(&req)

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorInternalDatabase": err})
		return
	}
	defer pool.Close()

	hashsedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorProcessingRequest": err})
	}

	myuuid := uuid.New()

	// your query to the database
	_, err = pool.Exec(context.Background(), "INSERT INTO users(id, username, pass, role) VALUES ($1, $2, $3, $4)", myuuid, req.Username, hashsedPassword, "client")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorProcessingRequest": err})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "signup complete"})
}
