package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	dir     string
	dirHTML string
	dirCSS  string
	dirTMPL string
)

var (
	r *gin.Engine
)

func init() {
	dir = os.Getenv("DIR_BS")
	dirHTML = dir + "/static/html"
	dirCSS = dir + "/static/css"
	dirTMPL = dir + "/tmpl"

	gin.ForceConsoleColor()
	r = gin.Default()
	r.SetTrustedProxies(nil)

	initReviews()
	initBookings()
	initCSS()
	initTemplates()
}

func Init() {
	r.GET("/", func(c *gin.Context) {
		c.File(dirHTML + "/index.html")
	})

	r.GET("/calendar.js", func(c *gin.Context) {
		c.File(dir + "/js/calendar.js")
	})

	r.Run(":5000")
}
