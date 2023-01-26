package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

var dirCSS string

func init() {
	dirCSS = os.Getenv("STATIC_DIR_BS") + "/css"
}

func initCSS(r *gin.Engine) {
	r.GET("/footer.css", func(c *gin.Context) {
		c.File(dirCSS + "/footer.css")
	})

	r.GET("/navbar.css", func(c *gin.Context) {
		c.File(dirCSS + "/navbar.css")
	})

	r.GET("/index.css", func(c *gin.Context) {
		c.File(dirCSS + "/index.css")
	})
}
