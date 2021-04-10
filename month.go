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
