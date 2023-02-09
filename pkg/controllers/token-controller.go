package controllers

import (
	"fmt"
	"net/http"

	"github.com/Allexsen/booking-site/pkg/models"
	"github.com/gin-gonic/gin"
)

// Controller and helper functions for tokens
// Used to validate/create/update tokens

func validateToken(token string) error {
	err := models.CheckToken(token)
	return err
}

func DisableToken(c *gin.Context) {
	token := c.Param("token")
	err := models.DisableToken(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.File(dirHTML + "/reviews")
}

func IsRewieavableToken(c *gin.Context) {
	token := c.Param("token")
	err := validateToken(token)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.Next()
}

func CreateToken(c *gin.Context) {
	token, err := models.CreateToken(c.PostForm("payment_id"), c.PostForm("booking_id"), c.PostForm("email"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// email should be sent using token
	// This is a placeholder for it
	fmt.Println(token)

	c.File(dirHTML + "/booking-success.html")
}
