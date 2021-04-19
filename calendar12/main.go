package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("/Users/andrewarrow/data/1900")
	s := string(b)
	on := false
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "moon_phase_name") {
			tokens := strings.Split(line, "<")
			for _, t := range tokens {
				if on {
					more := strings.Split(t, ">")
					if len(more) > 1 {
						fmt.Println(more[1])
					}
				}
				if t == "span>New Moon" || t == "span>Full Moon" {
					on = true
				}
				if t == "/tr>" {
					on = false
				}
			}
		}
	}
}

func main2() {
	year := 1901
	for {
		fmt.Printf("wget https://www.calendar-12.com/moon_phases/%d\n", year)
		fmt.Printf("sleep 1\n")
		year++
		if year > 2100 {
			break
		}
	}
}
