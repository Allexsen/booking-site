package controllers

import (
	"net/http"

	"github.com/Allexsen/booking-site/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// Controller and helper functions for reviews
// Used to process/handle requests regarding reviews
func LoadReviews(c *gin.Context) {
	if c.GetHeader("Accept") != "application/json" {
		c.File(dirHTML + "/reviews.html")
		return
	}

	r, err := models.GetReviews()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	reviews, err := json.Marshal(r)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, reviews)
}
