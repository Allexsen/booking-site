package routes

import (
	"github.com/Allexsen/booking-site/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func initBookings() {
	r.GET("/bookings", controllers.LoadBookings)

	r.POST("/bookings", controllers.HandleBooking)

	r.GET("/bookings/refund/:token", func(c *gin.Context) {
		c.String(200, "Refund hit")
	})
}
