package routes

import (
	"github.com/Allexsen/booking-site/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func initBookings(r *gin.Engine) {
	r.GET("/bookings", controllers.LoadBookings)

	r.POST("/bookings", func(c *gin.Context) {
		c.String(200, "Booking Payment hit")
	})

	r.GET("/bookings/refund/:token", func(c *gin.Context) {
		c.String(200, "Refund hit")
	})
}
