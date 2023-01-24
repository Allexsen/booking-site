package routes

import "github.com/gin-gonic/gin"

func Init() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	initReviews(r)
	initBookings(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Landing page hit")
	})

	r.Run(":5000")
}
