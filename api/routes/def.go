package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	initReviews(r)
	initBookings(r)
	initCSS(r)

	dir := os.Getenv("STATIC_DIR_BS")
	r.GET("/", func(c *gin.Context) {
		c.File(dir + "/html/index.html")
	})

	r.Run(":5000")
}
