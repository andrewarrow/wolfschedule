package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"html/template"

	"github.com/gin-gonic/gin"
)

var Times map[string]int64 = map[string]int64{}
var TimesMutex sync.Mutex

var Year int
var EventDate time.Time
var Month time.Month

func WelcomeIndex(c *gin.Context) {

	TimesMutex.Lock()
	body := template.HTML(makeHTML())
	TimesMutex.Unlock()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"body":  body,
	})
}

func makeHTML() string {
	buffer := []string{}
	for {
		count := 0
		buffer = append(buffer, fmt.Sprintf("<h2 class=\"mt-4\">%v %d</h2><div class=\"row mb-3\">\n", EventDate.Month(), Year))
		for {
			weekday := fmt.Sprintf("%v", EventDate.Weekday())
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
			u := fmt.Sprintf("%v", EventDate)
			substring := u[0:10]
			moon := ""
			if Times[substring] > 0 {
				moon = "moon"
			}
			buffer = append(buffer, fmt.Sprintf("<div class=\"col-2 themed-grid-col %s\">%d\n%s</div>", moon, EventDate.Day(), dayFinal))
			count++
			if count == 6 {
				count = 0
			}
			EventDate = EventDate.AddDate(0, 0, 1)
			if EventDate.Month() != Month {
				Month = EventDate.Month()
				buffer = append(buffer, "</div>")
				break
			}
		}

		if EventDate.Year() != Year {
			break
		}
	}

	return strings.Join(buffer, "\n")
}
