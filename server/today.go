package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/redis"
	"github.com/gin-gonic/gin"
)

func TodayIndex(c *gin.Context) {

	t := c.DefaultQuery("t", "0")
	tInt, _ := strconv.Atoi(t)

	tz, _ := c.Cookie("tz")
	if tz == "" {
		tz = "Antarctica/Troll"
	}

	body := template.HTML(makeTodayHTML(tInt, tz))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

type TimeZoneData struct {
	Zones []string
	Zone  string
}

func makeTodayHTML(offset int, tz string) string {
	buffer := []string{}

	location, _ := time.LoadLocation(tz)
	t := time.Now().In(location)
	t = t.Add(time.Hour)
	t = t.Add(time.Hour * time.Duration(offset))

	tmpl, err := template.ParseFiles("templates/tz.tmpl")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	b := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(b, TimeZoneData{zones, tz})
	if err != nil {
		fmt.Println(tz, err)
		return ""
	}

	buffer = append(buffer, "<p><h1>")
	buffer = append(buffer, t.Format("Monday Jan 01, 2006"))
	buffer = append(buffer, "</h1></p>")
	buffer = append(buffer, "<p>"+string(b.Bytes())+"</p>")
	buffer = append(buffer, "<p>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"?t=%d\">backwards</a> | <a href=\"?t=%d\">forward</a>", offset-1, offset+1))
	buffer = append(buffer, "</p>")
	buffer = append(buffer, "<div class=\"good-links\">")

	items := redis.QueryDay(offset)
	prevCount := 0
	for _, item := range items {
		if item.Count != prevCount {
			buffer = append(buffer, fmt.Sprintf("<h2>%d</h2>", item.Count))
		}
		buffer = append(buffer, fmt.Sprintf("<div class=\"item\"><a href=\"/item/%s\">%s</a></div>", url.QueryEscape(item.Title), item.Title))
		prevCount = item.Count
	}
	buffer = append(buffer, "</div>")

	return strings.Join(buffer, "\n")
}
