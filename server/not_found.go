package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotFoundIndex(c *gin.Context) {

	body := template.HTML(make404HTML())

	c.HTML(http.StatusNotFound, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func make404HTML() string {
	buffer := []string{}
	buffer = append(buffer, "<p><b>That's a wolf 404.</b></p>")
	return strings.Join(buffer, "\n")
}
