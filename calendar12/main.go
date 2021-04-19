package main

import "fmt"

func main() {
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
