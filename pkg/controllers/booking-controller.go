package controllers

import (
	"net/http"
	"os"

	"github.com/Allexsen/booking-site/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var dir string

func init() {
	dir = os.Getenv("HTML_DIR_BS")
}

// Controller and helper functions for bookings
// Used to process incoming booking requests
func LoadBookings(c *gin.Context) {
	if c.GetHeader("Accept") == "text/html" {
		c.File(dir + "/index.html")
		return
	}

	b, err := models.GetAllBookings()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	bookings, err := json.Marshal(b)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, bookings)
}
