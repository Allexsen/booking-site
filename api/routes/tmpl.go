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

// var (
// 	desc map[int]string
// 	msgs map[int]string
// )

// func init() {
// 	desc = make(map[int]string)
// 	desc[400] = "Bad Request!"
// 	desc[403] = "Forbidden"
// 	desc[404] = "Page not found!"

// 	msgs = make(map[int]string)
// 	msgs[400] = "Invalid request - can't be processed"
// 	msgs[403] = "Request not permitted"
// 	msgs[404] = "We are sorry, requested page couldn't be found."
// }

func initTemplates() {
	r.NoRoute(func(c *gin.Context) {
		tmpl := template.Must(template.ParseFiles(dirTMPL + "/err.html"))
		page := Page{
			Title:      "404 - Page not found",
			StatusCode: "404",
			Desc:       "Page not found",
			Message:    "We are sorry, requested page couldn't be found.",
		}

		err := tmpl.Execute(c.Writer, page)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})
}
