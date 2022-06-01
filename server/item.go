package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/wolfschedule/redis"
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

	m := redis.QueryAttributes(title)

	buffer = append(buffer, "<p><h1>")
	buffer = append(buffer, title)
	buffer = append(buffer, "</h1></p>")
	buffer = append(buffer, "<p>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"https://news.google.com%s\">source</a>", m["href"][1:]))
	for k, v := range m {
		if k == "href" {
			continue
		}
		buffer = append(buffer, "<div>")
		buffer = append(buffer, fmt.Sprintf("%s: %s", k, v))
		buffer = append(buffer, "</div>")
	}
	buffer = append(buffer, "</p>")
	return strings.Join(buffer, "\n")
}
