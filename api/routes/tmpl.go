package routes

import "github.com/gin-gonic/gin"

func initTemplates() {
	r.GET("/400", func(c *gin.Context) {
		c.String(200, "400 hit")
	})
}
