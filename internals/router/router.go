package router

import (
	"golang-movie-reservation/internals/services"
	"golang-movie-reservation/middleware"

	"github.com/gin-gonic/gin"
)

func Router(adminRoute *services.Admin, movieRoute *services.Global, clientRoute *services.Client) *gin.Engine {
	r := gin.Default()

	r.POST("/login", services.Login)
	r.POST("/signup", services.Signup)

	// admin only
	r.POST("/addmovie", middleware.AuthHandlerAdmin(), adminRoute.AddMovie)
	r.PUT("/editmovie", middleware.AuthHandlerAdmin(), adminRoute.EditMovie)
	r.DELETE("/deletemovie", middleware.AuthHandlerAdmin(), adminRoute.DeleteMovie)
	r.POST("/addschedule", middleware.AuthHandlerAdmin(), adminRoute.AddSchedule)
	r.PUT("/editschedule", middleware.AuthHandlerAdmin(), adminRoute.EditSchedule)
	r.DELETE("/deleteschedule", middleware.AuthHandlerAdmin(), adminRoute.DeleteSchedule)

	// both admin and client
	r.GET("/getmovie", movieRoute.GetMovie)

	// client
	r.POST("/reserve/cancel", middleware.AuthHandlerClient(), clientRoute.CancelReserve)
	r.POST("/reserve", middleware.AuthHandlerClient(), clientRoute.Reserve)

	// payment gateway
	r.POST("/payment", services.MakePayment)

	return r

}
