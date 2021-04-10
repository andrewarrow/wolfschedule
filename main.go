package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var special = "01/02/2006 15:04"

func PrintHelp() {
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "parse" {
		months := ParseData("2021.txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
	} else if command == "html" {
		months := ParseData("2021.txt")
		MakeHtml(months)
	}
}

func ParseData(f string) []Month {
	b, _ := ioutil.ReadFile(f)
	s := string(b)
	months := []Month{}
	ped := 0
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		newMoon := tokens[0] + " " + tokens[1]
		fullMoon := tokens[2] + " " + tokens[3]

		newDate, _ := time.Parse(special, newMoon)
		fullDate, _ := time.Parse(special, fullMoon)
		today := time.Now()

		m := handleMonth(int(newDate.Month()), newDate.Day(), fullDate.Day(), ped,
			int(today.Month()), today.Day())
		ped += m.EndDate
		months = append(months, m)
		//delta := fullDate.Unix() - newDate.Unix()
		//deltaString := fmt.Sprintf("%d", delta)
		//fmt.Println(newDate.Unix(), fullDate.Unix(), delta,
		//AsciiByteToBase9(deltaString))

	}
	return months
}

func handleMonth(m, d1, d2, ped, todayMonth, todayDay int) Month {
	day1, _ := time.Parse(special, fmt.Sprintf("%02d/01/2021 00:00", m))
	//moon1, _ := time.Parse(special, fmt.Sprintf("%02d/%02d/2021 00:00", m, d1))
	//moon2, _ := time.Parse(special, fmt.Sprintf("%02d/%02d/2021 00:00", m, d1))

	mm := Month{}
	mm.PrevEndDate = ped
	mm.Name = fmt.Sprintf("% 2d", day1.Month())
	for {
		day1 = day1.AddDate(0, 0, 1)
		if int(day1.Month()) != m {
			break
		}
		mm.EndDate = day1.Day()
		if int(day1.Month()) == todayMonth && day1.Day() == todayDay {
			mm.Today = todayDay
		}
	}
	mm.Event1 = d1
	mm.Event2 = d2
	return mm
}

func oldHandleMonth(m, d1, d2 int) {
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
			fmt.Printf("% 2d!!!!", day1.Day())
			fmt.Println("")
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
