package server

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	router := gin.Default()

	prefix := ""
	router.Static("/assets", prefix+"assets")
	router.GET("/", WelcomeIndex)

	AddTemplates(router, prefix)
	go router.Run(fmt.Sprintf(":%s", port))

	for {
		time.Sleep(time.Second)
	}

}

func AddTemplates(r *gin.Engine, prefix string) {
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob(prefix + "templates/*.tmpl")
}
