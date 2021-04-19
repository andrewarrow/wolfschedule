package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
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
	prevTick := int64(0)
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
			digit := AsciiByteToBase9(deltaString)
			tick := delta / 60
			if prevTick > 0 {
				fmt.Println(tokens[0], delta, digit, tick, tick-prevTick)
			}
			prevTick = tick
		}

		prevTime = ts
	}
	fmt.Println("")
	fmt.Println("")
}
func ParseForPattern4() {

	b, _ := ioutil.ReadFile("1970_2021.txt")
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) == 1 {
			continue
		}
		delta := tokens[4]
		deltaInt, _ := strconv.Atoi(delta)
		factor := tokens[6]
		factorInt, _ := strconv.Atoi(factor)
		tick := float64(deltaInt) / float64(factorInt)
		fmt.Println(tokens[0], tick/60)
	}
	fmt.Println("")
	fmt.Println("")
}
