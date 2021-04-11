package main

import (
	"fmt"
	"strings"
	"time"
)

type Month struct {
	Event1      int
	Event2      int
	Event3      int
	Name        string
	EndDate     int
	PrevEndDate int
	Today       int
}

func (m *Month) StringForToday(today time.Time) string {
	buff := []string{}

	day := 1
	for {
		if day == today.Day() {
			buff = append(buff, "_")
		}
		if m.Event1 == day {
			buff = append(buff, fmt.Sprintf("%02d", day))
			buff = append(buff, "\n")
		} else if m.Event2 == day {
			buff = append(buff, fmt.Sprintf("%02d", day))
			buff = append(buff, "\n")
		} else {
			buff = append(buff, fmt.Sprintf("%02d", day))
		}
		if day == today.Day() {
			buff = append(buff, "_")
		}
		if day == m.EndDate {
			break
		}
		day++
	}

	return strings.Join(buff, " ")
}
func (m *Month) String() string {
	buff := []string{}

	day := 1
	for {
		if m.Event1 == day {
			buff = append(buff, fmt.Sprintf("%02d ", day))
		} else if m.Event1-1 == day {
			buff = append(buff, " ")
		} else if m.Event1-2 == day {
			buff = append(buff, " ")
		} else if m.Event1+1 == day {
			buff = append(buff, " ")
		} else if m.Event1+2 == day {
			buff = append(buff, " ")
		} else if m.Event2 == day {
			buff = append(buff, fmt.Sprintf("%02d ", day))
		} else if m.Event2-1 == day {
			buff = append(buff, " ")
		} else if m.Event2-2 == day {
			buff = append(buff, " ")
		} else if m.Event2+1 == day {
			buff = append(buff, " ")
		} else if m.Event2+2 == day {
			buff = append(buff, " ")
		} else if m.Event3 == day {
			buff = append(buff, fmt.Sprintf("%02d ", day))
		} else if m.Event3-1 == day {
			buff = append(buff, " ")
		} else if m.Event3-2 == day {
			buff = append(buff, " ")
		} else if m.Event3 > 0 && m.Event3+1 == day {
			buff = append(buff, " ")
		} else if m.Event3 > 0 && m.Event3+2 == day {
			buff = append(buff, " ")
		} else {
			buff = append(buff, fmt.Sprintf("%02d ", day))
		}
		if day == m.EndDate {
			break
		}
		day++
	}

	return strings.Join(buff, "")
}
func (m *Month) HTML() string {
	buff := []string{}

	day := 1
	boxOpen := false
	for {
		if m.Event1 == day {
			buff = append(buff, fmt.Sprintf("<span style='color: red;'>%02d</span>&nbsp;", day))
		} else if m.Event2-1 == day {

			buff = append(buff, fmt.Sprintf("<span>%02d</span></span>&nbsp;", day))
		} else if m.Event2-2 == day {

			buff = append(buff, fmt.Sprintf("<span style='border: 1px solid black;'><span>%02d</span>&nbsp;", day))
		} else if m.Event1-1 == day {

			buff = append(buff, fmt.Sprintf("<span>%02d</span></span>&nbsp;", day))
		} else if m.Event1-2 == day {

			buff = append(buff, fmt.Sprintf("<span style='border: 1px solid black;'><span>%02d</span>&nbsp;", day))
		} else if m.Event1+1 == day {

			buff = append(buff, fmt.Sprintf("<span style='border: 1px solid black;'><span>%02d</span>&nbsp;", day))
		} else if m.Event1+2 == day {

			buff = append(buff, fmt.Sprintf("<span>%02d</span></span>&nbsp;", day))
		} else if m.Event2+1 == day {

			boxOpen = true
			buff = append(buff, fmt.Sprintf("<span style='border: 1px solid black;'><span>%02d</span>", day))
		} else if m.Event2+2 == day {

			buff = append(buff, fmt.Sprintf("&nbsp;<span>%02d</span></span>&nbsp;", day))
			boxOpen = false
		} else if m.Event2 == day {
			buff = append(buff, fmt.Sprintf("<span style='color: red;'>%02d</span>&nbsp;", day))
		} else {
			buff = append(buff, fmt.Sprintf("<span>%02d</span>&nbsp;", day))
		}
		if day == m.EndDate {
			if boxOpen {
				buff = append(buff, fmt.Sprintf("</span>"))
				boxOpen = false
			}
			break
		}
		day++
	}

	return strings.Join(buff, "")
}
