package routes

import (
	"github.com/Allexsen/booking-site/pkg/controllers"
)

func initBookings() {
	r.GET("/bookings", controllers.LoadBookings)

	r.POST("/bookings", controllers.HandleBooking)

	r.GET("/bookings/refund/:token", controllers.DisableToken)
}
