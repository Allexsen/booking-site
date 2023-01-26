package controllers

import (
	"net/http"
	"os"

	"github.com/Allexsen/booking-site/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// Controller and helper functions for bookings
// Used to process incoming booking requests
func LoadBookings(c *gin.Context) {
	b, err := models.GetAllBookings()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	bookings, err := json.Marshal(b)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	dir := os.Getenv("STATIC_DIR_BS") + "/html/bookings.html"
	c.File(dir)
	c.JSON(http.StatusOK, bookings)
}
