package controllers

import (
	"fmt"
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

// ToDo:
func LoadNewReviewPage(c *gin.Context) {
	c.File(dirHTML + "/review-new.html")
}

func AddReview(c *gin.Context) {
	rid, err := models.CreateReview(c.Param("rating"), c.Param("booking_id"), c.Param("body"), c.Param("author"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// This is a placeholder
	// Should instead show the new review on top of other reviews
	fmt.Println(rid)
	c.File(dirHTML + "/reviews")
}
