package server

import (
	"net/http"

	"html/template"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {

	body := template.HTML("<p>wefwefwef</p>")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}
