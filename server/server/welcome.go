package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
	})
}
