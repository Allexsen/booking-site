package routes

import "github.com/gin-gonic/gin"

func Init() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Landing page hit")
	})

	r.GET("/booking", func(c *gin.Context) {
		c.String(200, "Booking page hit")
	})

	r.GET("/reviews", func(c *gin.Context) {
		c.String(200, "Reviews page hit")
	})

	r.Run(":5000")
}
