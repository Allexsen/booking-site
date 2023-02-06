package controllers

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/Allexsen/booking-site/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var (
	dir     string
	dirHTML string
)

func init() {
	dir = os.Getenv("DIR_BS")
	dirHTML = dir + "/static/html"
}

// Controller and helper functions for bookings
// Used to process incoming booking requests
func LoadBookings(c *gin.Context) {
	if c.GetHeader("Accept") != "application/json" {
		c.File(dirHTML + "/bookings.html")
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

func HandleBooking(c *gin.Context) {
	stDate, err := time.Parse("2006-01-02", c.PostForm("start_date"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	endDate, err := time.Parse("2006-01-02", c.PostForm("end_date"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	if stDate.After(endDate) {
		stDate, endDate = endDate, stDate
	}

	b := models.Booking{
		StartDate: sql.NullTime{Time: stDate, Valid: true},
		EndDate:   sql.NullTime{Time: endDate, Valid: true},
	}

	err = b.Archive()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.File(dirHTML + "/booking-success.html")
}
