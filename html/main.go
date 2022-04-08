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
		fmt.Println(eventDate, eventDate.Weekday())
		eventDate = eventDate.AddDate(0, 0, 1)
		if eventDate.Month() != month {
			break
		}
	}

}
