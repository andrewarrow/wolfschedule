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
		months2021 := ParseData("2021.txt")
		months2022 := ParseData("2022.txt")
		both := append(months2021, months2022...)
		for _, m := range both {
			fmt.Println(m.Event1, m.Event2, m.Event3)
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
	b, _ := ioutil.ReadFile(f)
	s := string(b)
	months := []Month{}
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		var m Month
		if len(tokens) == 7 && tokens[6] == "blue" {
			newMoon := tokens[0] + " " + tokens[1]
			fullMoon := tokens[2] + " " + tokens[3]
			newMoon2 := tokens[4] + " " + tokens[5]
			newDate, _ := time.Parse(special, newMoon)
			fullDate, _ := time.Parse(special, fullMoon)
			newDate2, _ := time.Parse(special, newMoon2)
			m = handleMonth(&newDate, &fullDate, &newDate2)
		} else {
			newMoon := tokens[0] + " " + tokens[1]
			fullMoon := tokens[2] + " " + tokens[3]
			newDate, _ := time.Parse(special, newMoon)
			fullDate, _ := time.Parse(special, fullMoon)
			m = handleMonth(&newDate, &fullDate, nil)
		}
		months = append(months, m)
		//delta := fullDate.Unix() - newDate.Unix()
		//deltaString := fmt.Sprintf("%d", delta)
		//fmt.Println(newDate.Unix(), fullDate.Unix(), delta,
		//AsciiByteToBase9(deltaString))

	}
	return months
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

func handleMonth2(m, d1, d2, ped, todayMonth, todayDay int) Month {
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
