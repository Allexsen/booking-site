package routes

import "github.com/gin-gonic/gin"

func Init() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.File("/static/html/index.html")
	})
}
