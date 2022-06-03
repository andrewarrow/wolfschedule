package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"html/template"

	"github.com/andrewarrow/wolfschedule/moon"
	"github.com/gin-gonic/gin"
)

func YearIndex(c *gin.Context) {

	tInt, tz := TimeAndZone(c)
	location, _ := time.LoadLocation(tz)
	t := time.Unix(tInt, 0)
	t = t.In(location)

	events := moon.FindEventsForYear(t.Year(), location)
	body := template.HTML(makeYearHTML(t.Year(), location, events))

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeYearHTML(year int, location *time.Location, events []*moon.Event) string {
	jan1 := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, location)
	month := jan1.Month()

	buffer := []string{}
	for {
		count := 0
		buffer = append(buffer, fmt.Sprintf("<h2 class=\"mt-4\">%v %d</h2><div class=\"row mb-3\">\n", jan1.Month(), year))
		for {
			weekday := fmt.Sprintf("%v", jan1.Weekday())
			printDay := ""
			if weekday == "Monday" {
				printDay = "mon"
			} else if weekday == "Tuesday" {
				printDay = "tue"
			} else if weekday == "Wednesday" {
				printDay = "wed"
			} else if weekday == "Thursday" {
				printDay = "thu"
			} else if weekday == "Friday" {
				printDay = "fri"
			}
			dayFinal := ""
			if printDay != "" {
				dayFinal = fmt.Sprintf("<div class=\"day-of-week\">%s</div>", printDay)
			}
			//u := fmt.Sprintf("%v", jan1)
			//substring := u[0:10]
			moon := ""
			//if Times[substring] > 0 {
			//	moon = "moon"
			//}
			buffer = append(buffer, fmt.Sprintf("<div class=\"col-2 themed-grid-col %s\">%d\n%s</div>", moon, jan1.Day(), dayFinal))
			count++
			if count == 6 {
				count = 0
			}
			jan1 = jan1.AddDate(0, 0, 1)
			if jan1.Month() != month {
				month = jan1.Month()
				buffer = append(buffer, "</div>")
				break
			}
		}

		if jan1.Year() != year {
			break
		}
	}

	return strings.Join(buffer, "\n")
}
