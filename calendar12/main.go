package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	yearInt := 1969
	for {
		yearInt++
		if yearInt > 2100 {
			break
		}
		year := fmt.Sprintf("%d", yearInt)
		b, _ := ioutil.ReadFile("/Users/andrewarrow/data/" + year)
		s := string(b)
		on := false
		buff := []string{}
		for _, line := range strings.Split(s, "\n") {
			if strings.Contains(line, "moon_phase_name") {
				tokens := strings.Split(line, "<")
				for _, t := range tokens {
					if on {
						more := strings.Split(t, ">")
						if len(more) > 1 {
							if more[1] != "" {
								buff = append(buff, more[1])
							}
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

		buff2 := []string{}
		for i, line := range buff {
			if i%2 == 0 {
				item := line[0:(len(line)-3)] + year + " "
				buff2 = append(buff2, item)
			} else {
				buff2 = append(buff2, line+"")
			}
		}
		buff3 := []string{}
		for i := 0; i < len(buff2)-1; i += 2 {
			item := fmt.Sprintf("%s%s", buff2[i], buff2[i+1])
			buff3 = append(buff3, item)
		}
		for i, line := range buff3 {
			if i%2 != 0 {
				test, _ := time.Parse("January 2, 2006 15:04 PM", line)
				fmt.Println(test, test.Unix())
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
