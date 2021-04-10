package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	//ParseData("2020.txt")
	ParseData("2021.txt")
	//ParseData("2022.txt")
	//ParseData("demo.txt")
}

var special = "01/02/2006 15:04"

func ParseData(f string) {
	b, _ := ioutil.ReadFile(f)
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		newMoon := tokens[0] + " " + tokens[1]
		fullMoon := tokens[2] + " " + tokens[3]

		newDate, _ := time.Parse(special, newMoon)
		fullDate, _ := time.Parse(special, fullMoon)

		handleMonth(int(newDate.Month()), newDate.Day(), fullDate.Day())
		//delta := fullDate.Unix() - newDate.Unix()
		//deltaString := fmt.Sprintf("%d", delta)
		//fmt.Println(newDate.Unix(), fullDate.Unix(), delta,
		//AsciiByteToBase9(deltaString))

	}
	fmt.Println("")
}

func handleMonth(m, d1, d2 int) {
	day1, _ := time.Parse(special, fmt.Sprintf("%02d/01/2021 00:00", m))
	//moon1, _ := time.Parse(special, fmt.Sprintf("%02d/%02d/2021 00:00", m, d1))
	//moon2, _ := time.Parse(special, fmt.Sprintf("%02d/%02d/2021 00:00", m, d1))

	fmt.Printf("\n\n% 2d\n\n", day1.Month())
	for {
		if d1 == day1.Day() {
			fmt.Printf("% 2d!!!!", day1.Day())
			fmt.Println("")
		} else if d1-1 == day1.Day() {
			fmt.Printf("% 2d!!!", day1.Day())
		} else if d1-2 == day1.Day() {
			fmt.Printf("% 2d!!", day1.Day())
		} else if d1-3 == day1.Day() {
			fmt.Printf("% 2d!", day1.Day())
		} else if d2 == day1.Day() {
			fmt.Println("")
			fmt.Printf("% 2d!!!!", day1.Day())
		} else if d2-1 == day1.Day() {
			fmt.Printf("% 2d!!!", day1.Day())
		} else if d2-2 == day1.Day() {
			fmt.Printf("% 2d!!", day1.Day())
		} else if d2-3 == day1.Day() {
			fmt.Printf("% 2d!", day1.Day())
		} else {
			fmt.Printf("% 2d", day1.Day())
		}
		//if fmt.Sprintf("%v", day1.Weekday()) == "Saturday" {
		//	fmt.Println("")
		//}
		//fmt.Println(day1, day1.Weekday())
		day1 = day1.AddDate(0, 0, 1)
		if int(day1.Month()) != m {
			break
		}
	}
	fmt.Println("")

}
func AsciiByteToBase9(a string) byte {

	sum := byte(0)
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	if len(strSum) > 1 {
		return AsciiByteToBase9(strSum)
	}
	return sum

}
