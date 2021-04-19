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
	fmt.Println("By Default, it will display the current -15 +15 days")
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

func GetAll() []Thing {
	things = []Thing{}
	ParseData2("1970_2100.csv")
	sort.SliceStable(things, func(i, j int) bool {
		return things[i].Val < things[j].Val
	})
	return things
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
		DisplayCurrentDay(argMap["year"], 0)
		return
	}
	command := os.Args[1]

	if command == "parse" {
	} else if command == "images" {
	} else if strings.HasPrefix(command, "--year") {
		//MakeImages(myimage)
		DisplayCurrentDay(argMap["year"], 0)
	} else if strings.HasPrefix(command, "-f") {
		last, other := DisplayCurrentDay(argMap["year"], 0)
		fmt.Println("")
		delta := last - time.Now().Unix() // left to go
		days := float64(delta) / 86400
		seconds := delta % 86400
		distance := last - other
		done := distance - delta
		per := float64(done) / float64(distance)
		s := fmt.Sprintf("%0.2f day(s), %d second(s) %% %0.6f", days, seconds, per)
		fmt.Printf("Next Event in: %s", s)
		for {
			time.Sleep(time.Second * 1)
			delta = last - time.Now().Unix()
			days = float64(delta) / 86400
			seconds = delta % 86400
			done = distance - delta
			per = float64(done) / float64(distance)
			backspace := []byte{}
			for i := 0; i < len(s); i++ {
				backspace = append(backspace, 8)
			}
			fmt.Printf("%s", string(backspace))
			fmt.Printf("%0.2f day(s), %d second(s) %% %0.6f", days, seconds, per)
		}
	} else if strings.HasPrefix(command, "-") {
		add, _ := strconv.Atoi(command)
		DisplayCurrentDay(argMap["year"], add)
	} else if strings.HasPrefix(command, "+") {
		//MakeImages(myimage)
		add, _ := strconv.Atoi(command[1:])
		DisplayCurrentDay(argMap["year"], add)
	} else if command == "day" {
		DisplayCurrentDay(argMap["year"], 0)
	} else if command == "pdf" {
		MakePDF("1970", 1)
		MakePDF("1970", 2)
		MakePDF("1970", 3)
		MakePDF("1970", 4)
		MakePDF("1970", 5)
		MakePDF("1970", 6)
		MakePDF("1970", 7)
		MakePDF("1970", 8)
		MakePDF("1970", 9)
		MakePDF("1970", 10)
		MakePDF("1970", 11)
		MakePDF("1970", 12)
	} else if command == "earth" {
		EarthAge()
	} else if command == "pattern" {
		ParseForPattern()
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
	months = []Month{}

	b, _ := ioutil.ReadFile(f)
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, ",")
		if len(tokens) < 3 {
			break
		}
		ts, _ := strconv.ParseInt(tokens[1], 10, 64)
		eventDate := time.Unix(ts, 0)
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
