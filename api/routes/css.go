package routes

import (
	"github.com/gin-gonic/gin"
)

func initCSS() {
	r.GET("/css/:file", func(c *gin.Context) {
		c.File(dirCSS + "/" + c.Param("file"))
	})
}
