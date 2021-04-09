package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	b, _ := ioutil.ReadFile("2021.txt")
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		newMoon := tokens[0] + " " + tokens[1]
		fullMoon := tokens[2] + " " + tokens[3]

		newDate, _ := time.Parse("01/02/2006 15:04", newMoon)
		fullDate, _ := time.Parse("01/02/2006 15:04", fullMoon)

		fmt.Println(newDate, fullDate)
		fmt.Println(newDate.Unix(), fullDate.Unix(), fullDate.Unix()-newDate.Unix())

	}
}
