package routes

import "github.com/gin-gonic/gin"

func initBookings(r *gin.Engine) {
	r.GET("/bookings", func(c *gin.Context) {
		c.String(200, "Booking page hit")
	})

	r.POST("/bookings", func(c *gin.Context) {
		c.String(200, "Booking Payment hit")
	})

	r.GET("/bookings/refund/:token", func(c *gin.Context) {
		c.String(200, "Refund hit")
	})
}
