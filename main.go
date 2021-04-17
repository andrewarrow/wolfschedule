package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var special = "01/02/2006 15:04"
var curMonth = Month{}
var months = []Month{}
var things = []Thing{}

type Thing struct {
	Text string
	Val  int64
}

func PrintHelp() {
	fmt.Println("")
	fmt.Println("By Default, it will display the current year, full year view.")
	fmt.Println("")
	fmt.Println("wolfschedule help      # this menu")
	fmt.Println("wolfschedule today     # show me just enough for today")
	fmt.Println("wolfschedule           # --year=x")
	fmt.Println("wolfschedule next      # --offset=x")
	fmt.Println("wolfschedule prev      # --offset=x")
	fmt.Println("wolfschedule debug     # ")
	fmt.Println("wolfschedule wave      # display wave form")
	fmt.Println("")
}

func GetAll(y int) []Thing {
	things = []Thing{}
	for i := 2003; i < 2031; i++ {
		ParseData2(fmt.Sprintf("%d.txt", i))
	}
	sort.SliceStable(things, func(i, j int) bool {
		return things[i].Val < things[j].Val
	})
	all := []Thing{}
	for _, t := range things {
		u := time.Unix(t.Val, 0)
		if u.Year() == y-1 && int(u.Month()) == 12 {
			all = append(all, t)
		} else if u.Year() == y+1 && int(u.Month()) == 1 {
			all = append(all, t)
		} else if u.Year() == y {
			all = append(all, t)
		}
	}
	sort.SliceStable(all, func(i, j int) bool {
		return all[i].Val < all[j].Val
	})
	return all
	/*
		prev := int64(0)
		prevMonth := int(1)
		u := time.Now()
		delta := int64(0)
		deltas := []Delta{}
		times := []time.Time{}
		for _, t := range things {
			//fmt.Println(t.Val, t.Text)
			u = time.Unix(t.Val, 0)
			if int(u.Month()) != prevMonth {
				var m Month
				if len(times) == 3 {
					m = handleMonth(&times[0], &times[1], &times[2])
				} else if len(times) == 2 {
					m = handleMonth(&times[0], &times[1], nil)
				}
				months = append(months, m)
				//fmt.Println(prevMonth, times)
				times = []time.Time{}
			}
			times = append(times, u)
			if prev > 0 {
				delta = t.Val - prev
				d := NewDelta(int(delta), t.Text, int(u.Month()))
				d.Time = u
				deltas = append(deltas, d)
			}

			prev = t.Val
			prevMonth = int(u.Month())
		}
		//fmt.Println(prevMonth, times)
		var m Month
		if len(times) == 3 {
			m = handleMonth(&times[0], &times[1], &times[2])
		} else if len(times) == 2 {
			m = handleMonth(&times[0], &times[1], nil)
		}
		months = append(months, m)
	*/
}

func main() {
	rand.Seed(time.Now().UnixNano())
	argMap := argsToMap()

	if len(os.Args) == 1 {
		PrintHelp()
		months, _ := ParseData("2021.txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		fmt.Println("")
		return
	}
	command := os.Args[1]

	if argMap["year"] != "" && command != "wave" && command != "side" {
		months, _ := ParseData(argMap["year"] + ".txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		return
	}

	if command == "parse" {
	} else if command == "images" {
		//MakeImages(myimage)
	} else if command == "side" {
		y, _ := strconv.Atoi(argMap["year"])
		now := time.Now()
		today := fmt.Sprintf("%v", now)
		all := GetAll(y)
		m := map[string]bool{}
		for _, t := range all {
			u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
			m[u[0:10]] = true
		}
		day1 := time.Unix(all[0].Val, 0)
		last := time.Unix(all[len(all)-1].Val, 0)
		for {
			u := fmt.Sprintf("%v", day1)
			wd := fmt.Sprintf("%v", day1.Weekday())
			if wd == "Tuesday" || wd == "Thursday" || wd == "Saturday" ||
				wd == "Sunday" {
				wd = ""
			}
			substring := u[0:10]
			arrow := " "
			if substring == today[0:10] {
				arrow = " <---------------"
			}
			padding := ""
			otherPadding := ""
			if m[u[0:10]] {
				padding = ""
				otherPadding = "           "
			} else {
				otherPadding = ""
				padding = "           "
			}
			fmt.Printf("%s%s%s%s%s\n", padding, substring, otherPadding, arrow, fmt.Sprintf("%40s", wd))
			day1 = day1.AddDate(0, 0, 1)
			if day1.After(last) {
				break
			}
		}
		//MakeSides(deltas, fmt.Sprintf("%d", y))
	} else if command == "wave" {
		_, deltas := ParseData(argMap["year"] + ".txt")
		prevDays := 0.0
		isNext := false
		now := time.Now()
		//dir := ""
		for _, d := range deltas {
			if d.Time.After(now) {
				isNext = true
			}
			days := float64(d.Val) / 86400
			//dir = "down"
			if days > prevDays {
				//dir = "up"
			}
			if isNext && int(now.Month()) == d.Month {
				i := 0
				day1 := d.Time
				day1 = day1.AddDate(0, 0, -1)
				day1 = day1.AddDate(0, 0, -1)
				day1 = day1.AddDate(0, 0, -1)
				day1 = day1.AddDate(0, 0, -1)
				day1 = day1.AddDate(0, 0, -1)
				day1 = day1.AddDate(0, 0, -1)
				for {
					day1 = day1.AddDate(0, 0, 1)
					fmt.Printf("++++ %s %s\n", fmt.Sprintf("%v", day1)[:10], day1.Weekday())
					if i > 3 {
						break
					}
					i++
				}
			}
			digit := AsciiByteToBase9(fmt.Sprintf("%d", d.Val))
			if prevDays == 0.0 {
				fmt.Printf("%d %d, %.3f, %30s %s\n", digit, d.Val, days, d.Text, d.Time.Weekday())
			} else {
				fmt.Printf("%d %d, %.3f, %30s %.3f %s\n", digit, d.Val, days, d.Text, math.Abs(prevDays-days), d.Time.Weekday())
			}

			if isNext && int(now.Month()) == d.Month {
				day1 := d.Time
				i := 0
				for {
					day1 = day1.AddDate(0, 0, 1)
					fmt.Printf("---- %s %s\n", fmt.Sprintf("%v", day1)[:10], day1.Weekday())
					if i > 3 {
						break
					}
					i++
				}
				isNext = false
			}

			prevDays = days
		}
	} else if command == "debug" {
	} else if command == "next" {
	} else if command == "prev" {
	} else if command == "today" {
		months, _ := ParseData("2021.txt")
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
		months, _ := ParseData("2021.txt")
		MakeHtml(months)
	}
}

func ParseData2(f string) {
	timeZone, _ := time.LoadLocation("America/Phoenix")

	months = []Month{}

	b, _ := ioutil.ReadFile(f)
	s := string(b)
	monthInt := 0
	year := ""
	for _, line := range strings.Split(s, "\n") {
		//2022 February 16 16:59 UTC
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		year = tokens[0]
		yearInt, _ := strconv.Atoi(year)
		month := tokens[1]
		if month == "January" {
			monthInt = 1
		} else if month == "February" {
			monthInt = 2
		} else if month == "March" {
			monthInt = 3
		} else if month == "April" {
			monthInt = 4
		} else if month == "May" {
			monthInt = 5
		} else if month == "June" {
			monthInt = 6
		} else if month == "July" {
			monthInt = 7
		} else if month == "August" {
			monthInt = 8
		} else if month == "September" {
			monthInt = 9
		} else if month == "October" {
			monthInt = 10
		} else if month == "November" {
			monthInt = 11
		} else if month == "December" {
			monthInt = 12
		}
		day := tokens[2]
		dayInt, _ := strconv.Atoi(day)
		hourMin := tokens[3]
		tokens = strings.Split(hourMin, ":")
		hour := tokens[0]
		hourInt, _ := strconv.Atoi(hour)
		min := tokens[1]
		minInt, _ := strconv.Atoi(min)

		eventDate := time.Date(yearInt, time.Month(monthInt), dayInt, hourInt, minInt, 0, 0, timeZone)
		//fmt.Println(line, eventDate.Unix())
		thing := Thing{}
		thing.Text = line
		thing.Val = eventDate.Unix()
		things = append(things, thing)
	}
}
func ParseData(f string) ([]Month, []Delta) {
	timeZone, _ := time.LoadLocation("America/Phoenix")

	months = []Month{}

	b, _ := ioutil.ReadFile(f)
	s := string(b)
	monthInt := 0
	year := ""
	for _, line := range strings.Split(s, "\n") {
		//2022 February 16 16:59 UTC
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		year = tokens[0]
		yearInt, _ := strconv.Atoi(year)
		month := tokens[1]
		if month == "January" {
			monthInt = 1
		} else if month == "February" {
			monthInt = 2
		} else if month == "March" {
			monthInt = 3
		} else if month == "April" {
			monthInt = 4
		} else if month == "May" {
			monthInt = 5
		} else if month == "June" {
			monthInt = 6
		} else if month == "July" {
			monthInt = 7
		} else if month == "August" {
			monthInt = 8
		} else if month == "September" {
			monthInt = 9
		} else if month == "October" {
			monthInt = 10
		} else if month == "November" {
			monthInt = 11
		} else if month == "December" {
			monthInt = 12
		}
		day := tokens[2]
		dayInt, _ := strconv.Atoi(day)
		hourMin := tokens[3]
		tokens = strings.Split(hourMin, ":")
		hour := tokens[0]
		hourInt, _ := strconv.Atoi(hour)
		min := tokens[1]
		minInt, _ := strconv.Atoi(min)

		eventDate := time.Date(yearInt, time.Month(monthInt), dayInt, hourInt, minInt, 0, 0, timeZone)
		//fmt.Println(line, eventDate.Unix())
		thing := Thing{}
		thing.Text = line
		thing.Val = eventDate.Unix()
		things = append(things, thing)
	}
	sort.SliceStable(things, func(i, j int) bool {
		return things[i].Val < things[j].Val
	})
	prev := int64(0)
	prevMonth := int(1)
	u := time.Now()
	delta := int64(0)
	deltas := []Delta{}
	times := []time.Time{}
	for _, t := range things {
		//fmt.Println(t.Val, t.Text)
		u = time.Unix(t.Val, 0)
		if int(u.Month()) != prevMonth {
			var m Month
			if len(times) == 3 {
				m = handleMonth(&times[0], &times[1], &times[2])
			} else if len(times) == 2 {
				m = handleMonth(&times[0], &times[1], nil)
			}
			months = append(months, m)
			//fmt.Println(prevMonth, times)
			times = []time.Time{}
		}
		times = append(times, u)
		if prev > 0 {
			delta = t.Val - prev
			d := NewDelta(int(delta), t.Text, int(u.Month()))
			d.Time = u
			deltas = append(deltas, d)
		}

		prev = t.Val
		prevMonth = int(u.Month())
	}
	//fmt.Println(prevMonth, times)
	var m Month
	if len(times) == 3 {
		m = handleMonth(&times[0], &times[1], &times[2])
	} else if len(times) == 2 {
		m = handleMonth(&times[0], &times[1], nil)
	}
	months = append(months, m)
	return months, deltas
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
