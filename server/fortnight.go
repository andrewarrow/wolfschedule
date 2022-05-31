package server

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func FortnightIndex(c *gin.Context) {

	TimesMutex.Lock()
	body := template.HTML(makeFortnightHTML())
	TimesMutex.Unlock()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeFortnightHTML() string {
	buffer := []string{}
	return strings.Join(buffer, "\n")
}
