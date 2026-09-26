package services

import (
	"context"
	"golang-movie-reservation/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Global struct {
}

func NewGlobal() *Global {
	return &Global{}
}

type Movie struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Creator        string  `json:"creator"`
	Genre          string  `json:"genre"`
	ReleaseDate    string  `json:"release_date"`
	DescMovie      *string `json:"desc_movie"`
	PosterImageURL *string `json:"poster_image_url"`
	TicketPrice    float64 `json:"ticket_price"`
}

// get movies
func (d *Global) GetMovie(c *gin.Context) {
	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			creator,
			genre,
			release_date,
			desc_movie,
			poster_image_url,
			ticket_price
		FROM movies
		ORDER BY release_date DESC
		`,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	movies := make([]Movie, 0)

	for rows.Next() {
		var movie Movie

		err := rows.Scan(
			&movie.ID,
			&movie.Name,
			&movie.Creator,
			&movie.Genre,
			&movie.ReleaseDate,
			&movie.DescMovie,
			&movie.PosterImageURL,
			&movie.TicketPrice,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}
