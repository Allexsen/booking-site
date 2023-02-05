package routes

import (
	"github.com/gin-gonic/gin"
)

func initCSS() {
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
