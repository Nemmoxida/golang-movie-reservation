package main

import (
	"golang-movie-reservation/internals/router"
	"golang-movie-reservation/internals/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// initiate services objects
	adminRoute := services.NewAdmin()
	movieRoute := services.NewGlobal()
	clientRoute := services.NewClient()

	r := router.Router(adminRoute, movieRoute, clientRoute)

	r.Run(":3001")
}
