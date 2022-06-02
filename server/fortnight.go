package server

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/moon"
	"github.com/gin-gonic/gin"
)

func FortnightIndex(c *gin.Context) {

	body := template.HTML(makeFortnightHTML())

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeFortnightHTML() string {
	buffer := []string{}

	location, _ := time.LoadLocation("UTC")
	t := time.Now()
	event := moon.FindNextEvent(t.Unix())
	if event == nil || event.Prev == nil || event.Next == nil {
		buffer = append(buffer, "<p><b>END</b></p>")
		return strings.Join(buffer, "\n")
	}

	buffer = append(buffer, "<p><br/>")

	start := event.Prev.Timestamp
	formatStr := "Monday 2006-01-02"
	prev := event.Prev.AsTime(location).Format(formatStr)
	theEvent := event.AsTime(location).Format(formatStr)
	next := event.Next.AsTime(location).Format(formatStr)

	for {
		formatted := time.Unix(start, 0).In(location).Format(formatStr)
		if formatted == prev || formatted == theEvent || formatted == next {
			buffer = append(buffer, "<b>"+formatted+"</b><br/>")
		} else {
			buffer = append(buffer, formatted+"<br/>")
		}
		if start > event.Next.Timestamp {
			break
		}
		start += 86400
	}
	buffer = append(buffer, "</p>")

	return strings.Join(buffer, "\n")
}
