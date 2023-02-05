package routes

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Title      string
	StatusCode string
	Desc       string
	Message    string
}

var (
	desc map[int]string
	msgs map[int]string
)

func init() {
	desc[400] = "Bad Request!"
	desc[403] = "Forbidden"
	desc[404] = "Page not found!"

	msgs[400] = "Invalid request - can't be processed"
	msgs[403] = "Request not permitted"
	msgs[404] = "We are sorry, the requested page can't be found."
}

func initTemplates() {
	r.NoRoute(func(c *gin.Context) {
		tmpl := template.Must(template.ParseFiles(dirTMPL + "/err.html"))
		stCode := c.Request.Response.StatusCode
		page := Page{
			Title:      "Error" + desc[stCode],
			StatusCode: string(rune(stCode)),
			Desc:       desc[stCode],
			Message:    msgs[stCode],
		}

		err := tmpl.Execute(c.Writer, page)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})
}
