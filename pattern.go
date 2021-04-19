package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/otiai10/primes"
)

func ParseForPattern2() {

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
			fmt.Printf("%s", tokens[2])
		}

		prevYear = year
	}
	fmt.Println("")
	fmt.Println("")
}
func ParseForPattern() {

	b, _ := ioutil.ReadFile("1970_2100.csv")
	s := string(b)
	prevTime := int64(0)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, ",")
		if len(tokens) < 3 {
			break
		}
		ts, _ := strconv.ParseInt(tokens[1], 10, 64)
		if prevTime > 0 {
			delta := ts - prevTime
			//maybe := 84560
			//maybe := 90600
			//fmt.Println(delta, float64(delta)/86400, float64(delta)/float64(maybe))
			deltaString := fmt.Sprintf("%d", delta)
			factors := primes.Factorize(delta).All()
			digit := AsciiByteToBase9(deltaString)
			deltaString = fmt.Sprintf("%d", factors[len(factors)-1])
			fmt.Println(delta, digit, factors[len(factors)-1], AsciiByteToBase9(deltaString))
		}

		prevTime = ts
	}
	fmt.Println("")
	fmt.Println("")
}
