package main

import (
	"fmt"
	"time"
)

func main() {
	year := time.Now().Year()
	timeZone, _ := time.LoadLocation("America/Phoenix")
	eventDate := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, timeZone)
	month := eventDate.Month()

	for {
		count := 0
		fmt.Printf("<h2 class=\"mt-4\">%v %d</h2><div class=\"row mb-3\">\n", eventDate.Month(), year)
		for {
			weekday := fmt.Sprintf("%v", eventDate.Weekday())
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
			fmt.Printf("<div class=\"col-2 themed-grid-col\">%d\n%s</div>", eventDate.Day(), dayFinal)
			count++
			if count == 6 {
				count = 0
			}
			eventDate = eventDate.AddDate(0, 0, 1)
			if eventDate.Month() != month {
				month = eventDate.Month()
				fmt.Println("</div>")
				break
			}
		}

		if eventDate.Year() != year {
			break
		}
	}

}
