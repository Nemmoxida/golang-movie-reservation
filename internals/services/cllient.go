package services

import (
	"context"
	"golang-movie-reservation/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// reserve

type ReserveRequest struct {
	ShowtimeID string   `json:"showtime_id"`
	Seats      []string `json:"seats"`
	UserID     string   `json:"user_id"`
}

func (d *Client) Reserve(c *gin.Context) {
	var req ReserveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(req.Seats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "at least one seat is required",
		})
		return
	}

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer pool.Close()

	ctx := context.Background()

	// transanction start
	tx, err := pool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer tx.Rollback(ctx)

	// check if the requested seats are avaiable
	for _, seat := range req.Seats {
		var exists bool

		err := tx.QueryRow(
			ctx,
			`
			SELECT EXISTS (
				SELECT 1
				FROM seats
				WHERE showtime_id = $1
				AND seat = $2
			)
			`,
			req.ShowtimeID,
			seat,
		).Scan(&exists)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if exists {
			c.JSON(http.StatusConflict, gin.H{
				"error": "seat " + seat + " is already reserved",
			})
			return
		}

		// add seat to query
		seatID := uuid.New().String()

		_, err = tx.Exec(
			ctx,
			`
			INSERT INTO seats (
				id,
				showtime_id,
				seat,
				status,
				user_id
			)
			VALUES ($1, $2, $3, $4, $5)
			`,
			seatID,
			req.ShowtimeID,
			seat,
			"reserved",
			req.UserID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	// check if transanction err
	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "reservation successfully created",
		"showtime": req.ShowtimeID,
		"seats":    req.Seats,
	})
}

// cancel reserve

type CancelReserveRequest struct {
	ShowtimeID string `json:"showtime_id"`
	Seat       string `json:"seat"`
	UserID     string `json:"user_id"`
}

func (d *Client) CancelReserve(c *gin.Context) {
	var req CancelReserveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer pool.Close()

	result, err := pool.Exec(
		context.Background(),
		`
		DELETE FROM seats
		WHERE showtime_id = $1
		AND seat = $2
		AND user_id = $3
		`,
		req.ShowtimeID,
		req.Seat,
		req.UserID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "reservation not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "reservation successfully cancelled",
	})
}
