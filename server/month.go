package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MonthIndex(c *gin.Context) {

	TimesMutex.Lock()
	body := template.HTML(makeMonthHTML())
	TimesMutex.Unlock()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeMonthHTML() string {
	buffer := []string{}
	return strings.Join(buffer, "\n")
}
