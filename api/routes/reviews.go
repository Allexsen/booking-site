package routes

import "github.com/gin-gonic/gin"

func initReviews(r *gin.Engine) {
	r.GET("/reviews", func(c *gin.Context) {
		c.String(200, "Reviews page hit")
	})

	r.GET("/reviews/add/:token", func(c *gin.Context) {
		c.String(200, "New Review GET hit on "+c.Param("token"))
	})

	r.POST("/reviews/add/:token", func(c *gin.Context) {
		c.String(200, "New Review POST hit on "+c.Param("token"))
	})
}
