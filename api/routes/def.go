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

	r.GET("/", func(c *gin.Context) {
		c.File(os.Getenv("STATIC_DIR_BS") + "/html/index.html")
	})

	r.Run(":5000")
}
