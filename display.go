package main

import (
	"fmt"
	"time"
)

func DisplayCurrentDay(year string, add int) {
	//y, _ := strconv.Atoi(year)
	now := time.Now()
	if year == "" {
		//	y = now.Year()
	}
	now = now.AddDate(0, 0, add)
	today := fmt.Sprintf("%v", now)
	all := GetAll()
	m := map[string]bool{}
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = true
	}
	day1 := now.AddDate(0, 0, -15)
	b1 := now.AddDate(0, 0, +15)
	for {
		u := fmt.Sprintf("%v", day1)
		wd := fmt.Sprintf("%v", day1.Weekday())
		if wd == "Tuesday" || wd == "Thursday" || wd == "Saturday" ||
			wd == "Sunday" {
			wd = ""
		}

		col1 := "" // event date
		col2 := "" // normal date
		col3 := "" // arrow
		col4 := "" // wd

		substring := u[0:10]
		if m[substring] {
			col1 = substring
		} else {
			col2 = substring
		}
		if substring == today[0:10] {
			col3 = " <-------------"
		}
		col4 = wd
		fmt.Printf("%10s %10s%20s%30s\n", col1, col2, col3, col4)
		day1 = day1.AddDate(0, 0, 1)
		if day1.After(b1) {
			break
		}
	}
}
