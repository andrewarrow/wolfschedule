package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ParseForPattern() {

	b, _ := ioutil.ReadFile("1970_2100.csv")
	s := string(b)
	year := 0
	prevYear := 0
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, ",")
		if len(tokens) < 3 {
			break
		}
		ts, _ := strconv.ParseInt(tokens[1], 10, 64)
		eventDate := time.Unix(ts, 0)
		year = eventDate.Year()
		if prevYear > 0 && year != prevYear {
			fmt.Println("")
		} else {
			fmt.Printf("%s ", tokens[2])
		}

		prevYear = year
	}
}
