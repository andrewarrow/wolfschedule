package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/redis"
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
	buffer = append(buffer, "<div>")

	items := redis.QueryDay()
	for _, item := range items {
		buffer = append(buffer, fmt.Sprintf("<div>%02d. %s</div>", item.Count, item.Title))
	}
	buffer = append(buffer, "</div>")

	return strings.Join(buffer, "\n")
}
