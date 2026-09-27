package services

import (
	"context"
	"golang-movie-reservation/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// movie
type AddMovieRequest struct {
	Name           string    `json:"name"`
	Creator        string    `json:"creator"`
	Genre          []string  `json:"genre"`
	ReleaseDate    time.Time `json:"release_date"`
	DescMovie      *string   `json:"desc_movie"`
	PosterImageURL *string   `json:"poster_image_url"`
	TicketPrice    float64   `json:"ticket_price"`
}

type EditMovieRequest struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Creator        string    `json:"creator"`
	Genre          string    `json:"genre"`
	ReleaseDate    time.Time `json:"release_date"`
	DescMovie      *string   `json:"desc_movie"`
	PosterImageURL *string   `json:"poster_image_url"`
	TicketPrice    float64   `json:"ticket_price"`
}

func (d *Admin) AddMovie(c *gin.Context) {
	var req AddMovieRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error finding request body",
			"error":   err.Error(),
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

	id := uuid.New().String()

	query := `
		INSERT INTO movies (
			id,
			name,
			creator,
			genre,
			release_date,
			desc_movie,
			poster_image_url,
			ticket_price
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = pool.Exec(
		context.Background(),
		query,
		id,
		req.Name,
		req.Creator,
		req.Genre,
		req.ReleaseDate,
		req.DescMovie,
		req.PosterImageURL,
		req.TicketPrice,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "movie successfully added",
		"id":      id,
	})
}

func (d *Admin) EditMovie(c *gin.Context) {
	var req EditMovieRequest

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

	query := `
		UPDATE movies
		SET
			name = $1,
			creator = $2,
			genre = $3,
			release_date = $4,
			desc_movie = $5,
			poster_image_url = $6,
			ticket_price = $7
		WHERE id = $8
	`

	result, err := pool.Exec(
		context.Background(),
		query,
		req.Name,
		req.Creator,
		req.Genre,
		req.ReleaseDate,
		req.DescMovie,
		req.PosterImageURL,
		req.TicketPrice,
		req.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "movie successfully updated",
	})
}

func (d *Admin) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer pool.Close()

	query := `
		DELETE FROM movies
		WHERE id = $1
	`

	result, err := pool.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "movie successfully deleted",
	})
}

// schedule
type AddScheduleRequest struct {
	MovieID   string    `json:"movie_id"`
	AddBy     string    `json:"add_by"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}

type EditScheduleRequest struct {
	ID        string    `json:"id"`
	MovieID   string    `json:"movie_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}

func (d *Admin) AddSchedule(c *gin.Context) {
	var req AddScheduleRequest

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

	id := uuid.New().String()

	query := `
		INSERT INTO showtimes (
			id,
			movie_id,
			add_by,
			date_start,
			date_end
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = pool.Exec(
		context.Background(),
		query,
		id,
		req.MovieID,
		req.AddBy,
		req.DateStart,
		req.DateEnd,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "schedule successfully added",
		"id":      id,
	})
}

func (d *Admin) EditSchedule(c *gin.Context) {
	var req EditScheduleRequest

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

	query := `
		UPDATE showtimes
		SET
			movie_id = $1,
			date_start = $2,
			date_end = $3
		WHERE id = $4
	`

	result, err := pool.Exec(
		context.Background(),
		query,
		req.MovieID,
		req.DateStart,
		req.DateEnd,
		req.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "schedule not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "schedule successfully updated",
	})
}

func (d *Admin) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer pool.Close()

	query := `
		DELETE FROM showtimes
		WHERE id = $1
	`

	result, err := pool.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "schedule not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "schedule successfully deleted",
	})
}
