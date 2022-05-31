package server

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func TodayIndex(c *gin.Context) {

	TimesMutex.Lock()
	body := template.HTML(makeTodayHTML())
	TimesMutex.Unlock()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeTodayHTML() string {
	buffer := []string{}

	buffer = append(buffer, "<p><h1>")
	buffer = append(buffer, time.Now().String()[0:36])
	buffer = append(buffer, "</h1></p>")

	return strings.Join(buffer, "\n")
}
