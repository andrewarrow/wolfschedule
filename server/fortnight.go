package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/moon"
	"github.com/gin-gonic/gin"
)

func FortnightIndex(c *gin.Context) {

	t := c.DefaultQuery("t", fmt.Sprintf("%d", time.Now().Unix()))
	tInt, _ := strconv.ParseInt(t, 10, 64)
	body := template.HTML(makeFortnightHTML(tInt))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeFortnightHTML(current int64) string {
	buffer := []string{}

	location, _ := time.LoadLocation("UTC")
	t := time.Unix(current, 0)
	event := moon.FindNextEvent(t.Unix())
	if event == nil || event.Prev == nil || event.Next == nil {
		buffer = append(buffer, "<p><b>END</b></p>")
		return strings.Join(buffer, "\n")
	}

	buffer = append(buffer, "<p>")

	start := event.Prev.Timestamp
	formatStr := "Monday 2006-01-02"
	now := t.In(location).Format(formatStr)
	prev := event.Prev.AsTime(location).Format(formatStr)
	theEvent := event.AsTime(location).Format(formatStr)
	next := event.Next.AsTime(location).Format(formatStr)

	for {
		formatted := time.Unix(start, 0).In(location).Format(formatStr)
		if formatted == prev {
			buffer = append(buffer, fmt.Sprintf("<div class=\"item\"><a href=\"?t=%d\"><b>%s</b></a></div>", event.Prev.Timestamp-1, formatted))
		} else if formatted == theEvent {
			buffer = append(buffer, fmt.Sprintf("<div class=\"item\"><a href=\"?t=%d\"><b>%s</b></a></div>", event.Timestamp, formatted))
		} else if formatted == next {
			buffer = append(buffer, fmt.Sprintf("<div class=\"item\"><a href=\"?t=%d\"><b>%s</b></a></div>", event.Next.Timestamp, formatted))
		} else if formatted == now {
			buffer = append(buffer, "<div class=\"item today\">"+formatted+"</div>")
		} else {
			buffer = append(buffer, "<div class=\"item\">"+formatted+"</div>")
		}
		if start > event.Next.Timestamp {
			break
		}
		start += 86400
	}
	buffer = append(buffer, "</p>")

	return strings.Join(buffer, "\n")
}
