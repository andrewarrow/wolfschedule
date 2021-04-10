package main

import (
	"fmt"
	"strings"
)

type Month struct {
	Event1      int
	Event2      int
	Name        string
	EndDate     int
	PrevEndDate int
	Today       int
}

func (m *Month) String() string {
	buff := []string{}

	day := 1
	for {
		if m.Event1 == day {
			buff = append(buff, "  ")
		} else if m.Event2 == day {
			buff = append(buff, "  ")
		} else {
			buff = append(buff, fmt.Sprintf("%d ", day))
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
	for {
		if m.Event1 == day {
			buff = append(buff, fmt.Sprintf("<span style='color: red;'>%02d</span>&nbsp;", day))
		} else if m.Event2-1 == day {

			buff = append(buff, fmt.Sprintf("<span>%02d</span></span>&nbsp;", day))
		} else if m.Event2-2 == day {

			buff = append(buff, fmt.Sprintf("<span style='border: 1px solid black;'><span>%02d</span>&nbsp;", day))
		} else if m.Event2+1 == day {

			buff = append(buff, fmt.Sprintf("<span style='color: cyan;'>%02d</span>&nbsp;", day))
		} else if m.Event2+2 == day {

			buff = append(buff, fmt.Sprintf("<span style='color: cyan;'>%02d</span>&nbsp;", day))
		} else if m.Event2 == day {
			buff = append(buff, fmt.Sprintf("<span style='color: red;'>%02d</span>&nbsp;", day))
		} else {
			buff = append(buff, fmt.Sprintf("<span>%02d</span>&nbsp;", day))
		}
		if day == m.EndDate {
			break
		}
		day++
	}

	return strings.Join(buff, "")
}
