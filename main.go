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
var curMonth = Month{}
var months = []Month{}

func PrintHelp() {
	fmt.Println("")
	fmt.Println("By Default, it will display the current year, full year view.")
	fmt.Println("")
	fmt.Println("wolfschedule help      # this menu")
	fmt.Println("wolfschedule today     # show me just enough for today")
	fmt.Println("wolfschedule           # --year=x")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	argMap := argsToMap()

	if len(os.Args) == 1 {
		PrintHelp()
		months := ParseData("2021.txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		fmt.Println("")
		return
	}
	command := os.Args[1]

	if argMap["year"] != "" {
		months := ParseData(argMap["year"] + ".txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		return
	}

	if command == "parse" {
		months2020 := ParseData("2020.txt")
		months2021 := ParseData("2021.txt")
		months2022 := ParseData("2022.txt")
		months2023 := ParseData("2023.txt")
		all := []Month{}
		all = append(all, months2020...)
		all = append(all, months2021...)
		all = append(all, months2022...)
		all = append(all, months2023...)
		times := []int64{}
		for _, m := range all {
			fmt.Println(m.Event1Unix, m.Event2Unix, m.Event3Unix)
			times = append([]int64{m.Event1Unix}, times...)
			times = append([]int64{m.Event2Unix}, times...)
			if m.Event3Unix > 0 {
				times = append([]int64{m.Event3Unix}, times...)
			}
		}
		prevTime := int64(0)
		for _, t := range times {
			if prevTime > 0 {
				delta := t - prevTime
				if delta < 0 {
					delta = delta * -1
				}
				deltaString := fmt.Sprintf("%d", delta)
				digit := AsciiByteToBase9(deltaString)
				fmt.Println(digit)
			}
			prevTime = t
		}
	} else if command == "today" {
		months := ParseData("2021.txt")
		today := time.Now()
		fmt.Println(today.Month())
		for _, m := range months {
			if fmt.Sprintf("% 2d", today.Month()) != m.Name {
				continue
			}
			fmt.Println(m.StringForToday(today))
		}
	} else if command == "help" {
		PrintHelp()
	} else if command == "html" {
		months := ParseData("2021.txt")
		MakeHtml(months)
	}
}

func ParseData(f string) []Month {
	timeZone, _ := time.LoadLocation("America/Phoenix")

	b, _ := ioutil.ReadFile(f)
	s := string(b)
	monthInt := 0
	year := ""
	bigMap := map[string][]string{}
	for _, line := range strings.Split(s, "\n") {
		//2022 February 16 16:59 UTC
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		month := tokens[1]
		bigMap[month] = append(bigMap[month], line)
	}
	text := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	for i, month := range text {
		curMonth = Month{}
		for j, line := range bigMap[month] {
			monthInt = i + 1
			tokens := strings.Split(line, " ")
			year = tokens[0]
			yearInt, _ := strconv.Atoi(year)
			day := tokens[2]
			dayInt, _ := strconv.Atoi(day)
			hourMin := tokens[3]
			tokens = strings.Split(hourMin, ":")
			hour := tokens[0]
			hourInt, _ := strconv.Atoi(hour)
			min := tokens[1]
			minInt, _ := strconv.Atoi(min)

			eventDate := time.Date(yearInt, time.Month(monthInt), dayInt, hourInt, minInt, 0, 0, timeZone)
			if j == 0 {
				curMonth.Event1 = dayInt
				curMonth.Event1Unix = eventDate.Unix()
			} else if j == 1 {
				curMonth.Event2 = dayInt
				curMonth.Event2Unix = eventDate.Unix()
			} else if j == 2 {
				curMonth.Event3 = dayInt
				curMonth.Event3Unix = eventDate.Unix()
			}
		}
		MakeDaysAnd(monthInt, year)
	}
	return months
}

func MakeDaysAnd(monthInt int, year string) {
	day1, _ := time.Parse(special, fmt.Sprintf("%02d/01/%s 00:00", monthInt, year))
	curMonth.Name = fmt.Sprintf("% 2d", day1.Month())
	for {
		day1 = day1.AddDate(0, 0, 1)
		if int(day1.Month()) != monthInt {
			break
		}
		curMonth.EndDate = day1.Day()
	}
	months = append(months, curMonth)
}

func handleMonth(newDate, fullDate, newDate2 *time.Time) Month {
	mm := Month{}
	m := int(newDate.Month())
	mm.Event1 = newDate.Day()
	mm.Event1Unix = newDate.Unix()
	mm.Event2 = fullDate.Day()
	mm.Event2Unix = fullDate.Unix()
	if newDate2 != nil {
		mm.Event3 = newDate2.Day()
		mm.Event3Unix = newDate2.Unix()
	}
	day1, _ := time.Parse(special, fmt.Sprintf("%02d/01/2021 00:00", m))
	mm.Name = fmt.Sprintf("% 2d", day1.Month())
	for {
		day1 = day1.AddDate(0, 0, 1)
		if int(day1.Month()) != m {
			break
		}
		mm.EndDate = day1.Day()
	}
	return mm
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
