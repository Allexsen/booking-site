package routes

import "github.com/gin-gonic/gin"

func Init() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Landing Page Hit")
	})

	r.Run(":5000")
}
