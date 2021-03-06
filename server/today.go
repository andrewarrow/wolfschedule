package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/moon"
	"github.com/gin-gonic/gin"
)

type TimeZoneData struct {
	Zones []string
	Zone  string
}
type ChartData struct {
	Data   []int
	Labels []string
}

func TodayIndex(c *gin.Context) {

	tInt, tz := TimeAndZone(c)

	body := template.HTML(makeTodayHTML(tInt, tz))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeTodayHTML(current int64, tz string) string {
	buffer := []string{}

	location, _ := time.LoadLocation(tz)
	t := time.Unix(current, 0).In(location)
	//t = t.Add(time.Hour * 24 * time.Duration(offset))

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
	buffer = append(buffer, t.Format(time.RFC850))
	buffer = append(buffer, "</h1></p>")
	buffer = append(buffer, "<p>"+string(b.Bytes())+"</p>")

	event := moon.FindNextEvent(t.Unix())
	if event == nil || event.Prev == nil {
		buffer = append(buffer, "<p><b>END</b></p>")
		return strings.Join(buffer, "\n")
	}
	if event.FullMoon {
		buffer = append(buffer, "<p><b>Next FULL MOON in</b></p>")
	} else {
		buffer = append(buffer, "<p><b>Next NEW MOON in</b></p>")
	}
	buffer = append(buffer, "<p>")
	buffer = append(buffer, moon.EventDelta(event.Timestamp-t.Unix())+"<br/>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"?t=%d\">%s</a>", event.Timestamp, event.AsTime(location).Format(time.RFC850)))
	buffer = append(buffer, "</p>")

	if event.FullMoon {
		buffer = append(buffer, "<p><b>Previous NEW MOON was</b></p>")
	} else {
		buffer = append(buffer, "<p><b>Previous FULL MOON was</b></p>")
	}
	buffer = append(buffer, "<p>")
	buffer = append(buffer, moon.EventDelta(t.Unix()-event.Prev.Timestamp)+" ago<br/>")
	buffer = append(buffer, fmt.Sprintf("<a href=\"?t=%d\">%s</a>", event.Prev.Timestamp, event.Prev.AsTime(location).Format(time.RFC850)))
	buffer = append(buffer, "</p>")

	tmpl, _ = template.ParseFiles("templates/chart.tmpl")
	b = bytes.NewBuffer([]byte{})
	cd := moon.FindEventsForChart()
	tmpl.Execute(b, ChartData{cd.Deltas, cd.Labels})
	buffer = append(buffer, "<div>"+string(b.Bytes())+"</div>")

	return strings.Join(buffer, "\n")
}
