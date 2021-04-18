package main

import (
	"fmt"
	"time"
)

func DisplayCurrentDay(year string, add int) int64 {
	//y, _ := strconv.Atoi(year)
	now := time.Now()
	if year == "" {
		//	y = now.Year()
	}
	now = now.AddDate(0, 0, add)
	today := fmt.Sprintf("%v", now)
	all := GetAll()
	m := map[string]int64{}
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = t.Val
	}
	day1 := now.AddDate(0, 0, -25)
	b1 := day1.AddDate(0, 0, +38)
	last := int64(0)
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
		if m[substring] > 0 {
			col1 = substring
			last = m[substring]
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
	return last
}

func MakePDF(year string, month int) {
	all := GetAll()
	m := map[string]string{}
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = "0"
	}
	day1, _ := time.Parse(special, fmt.Sprintf("01/01/2003 00:00"))
	day365, _ := time.Parse(special, fmt.Sprintf("12/31/2030 00:00"))
	eventHappened := -1
	for {
		u := fmt.Sprintf("%v", day1)
		substring := u[0:10]
		if m[substring] == "0" {
			eventHappened = 0
			inner := day1
			inner = inner.AddDate(0, 0, -5)
			for i := 4; i >= 0; i-- {
				innerU := fmt.Sprintf("%v", inner)
				innerSub := innerU[0:10]
				if m[innerSub] == "?" {
					m[innerSub] = fmt.Sprintf("%d", i+1)
				}
				inner = inner.AddDate(0, 0, 1)
			}
		} else if eventHappened <= 5 && eventHappened > 0 {
			m[substring] = fmt.Sprintf("%d", eventHappened)
		} else if eventHappened > 5 && eventHappened > 0 {
			m[substring] = "?"
		}
		eventHappened++
		day1 = day1.AddDate(0, 0, 1)
		if day1.After(day365) {
			break
		}
	}
	fmt.Println(m)
	b1 := day365.AddDate(0, 0, +5)
	hits := 0
	afterHit := 0

	buff := []string{}
	day1, _ = time.Parse(special, fmt.Sprintf("01/01/%s 00:00", year))
	day1 = day1.AddDate(0, 0, -5)
	day365, _ = time.Parse(special, fmt.Sprintf("12/31/%s 00:00", year))
	for {
		u := fmt.Sprintf("%v", day1)
		wd := fmt.Sprintf("%v", day1.Weekday())
		if wd == "Tuesday" || wd == "Thursday" || wd == "Saturday" ||
			wd == "Sunday" {
			wd = ""
		}

		col1 := "" // event date
		col2 := "" // normal date
		col4 := "" // wd

		substring := u[0:10]
		if m[substring] == "0" {
			col1 = substring
			hits++
			afterHit = 0
		} else {
			col2 = substring
		}
		col4 = wd
		thing := fmt.Sprintf("%10s %10s%20s%30s", col1, col2, "", col4)
		buff = append(buff, thing)
		if month == 1 && (hits == 0 || hits == 1 || hits == 2 || hits == 3) {
			fmt.Println(thing)
		} else if month == 2 && (hits == 3 || hits == 5 || hits == 6 || hits == 7) {
			fmt.Println(thing)
		}
		afterHit++

		//if hits == 3 && afterHit == 4 {
		//	break
		//}
		day1 = day1.AddDate(0, 0, 1)
		if day1.After(b1) {
			break
		}
	}
}
