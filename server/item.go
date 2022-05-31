package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ItemIndex(c *gin.Context) {

	title := c.Param("title")
	title = strings.Replace(title, "+", " ", -1)
	body := template.HTML(makeItemHTML(title))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeItemHTML(title string) string {
	buffer := []string{}

	buffer = append(buffer, "<p><h1>")
	buffer = append(buffer, title)
	buffer = append(buffer, "</h1></p>")
	return strings.Join(buffer, "\n")
}
