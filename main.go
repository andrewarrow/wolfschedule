package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
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

func GetAll() []Month {
	all := []Month{}
	for i := 2003; i < 2031; i++ {
		months, _ := ParseData(fmt.Sprintf("%d.txt", i))
		all = append(all, months...)
	}
	sort.SliceStable(things, func(i, j int) bool {
		return things[i].Val < things[j].Val
	})
	return all
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

	if argMap["year"] != "" && command != "wave" {
		months, _ := ParseData(argMap["year"] + ".txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		return
	}

	if command == "parse" {
		times := []float64{}
		for _, m := range GetAll() {
			//fmt.Println(m.Event1Unix, m.Event2Unix, m.Event3Unix)
			if m.Event1Unix > 0 {
				times = append([]float64{float64(m.Event1Unix)}, times...)
			}
			if m.Event2Unix > 0 {
				times = append([]float64{float64(m.Event2Unix)}, times...)
			}
			if m.Event3Unix > 0 {
				times = append([]float64{float64(m.Event3Unix)}, times...)
			}
		}
		sort.Float64s(times)
		prevTime := int64(0)
		for _, t := range times {
			if prevTime > 0 {
				delta := int64(t) - prevTime
				//fmt.Println(int64(t), prevTime, delta)
				if delta < 0 {
					delta = prevTime - int64(t)
					//fmt.Println("-1", t, prevTime, delta)
				}
				fmt.Println(delta / 3600)
				//deltaString := fmt.Sprintf("%d", delta)
				//digit := AsciiByteToBase9(deltaString)
				//fmt.Println(deltaString)
			}
			prevTime = int64(t)
		}
	} else if command == "wave" {
		months, deltas := ParseData(argMap["year"] + ".txt")
		for _, m := range months {
			fmt.Println(m.String())
		}
		fmt.Println(deltas)
	} else if command == "debug" {
		all := GetAll()
		encList := []string{}
		prev := int64(0)
		//prevDigit := byte(0)
		for _, t := range things {
			digit := byte(0)
			if prev > 0 {
				delta := t.Val - prev
				deltaString := fmt.Sprintf("%d", delta)
				last := deltaString[1 : len(deltaString)-1]
				lastInt, _ := strconv.Atoi(last)

				bs := make([]byte, 4)
				binary.LittleEndian.PutUint32(bs, uint32(lastInt))
				//enc := b64.StdEncoding.EncodeToString(bs)

				digit = AsciiByteToBase9(deltaString)
				encList = append(encList, fmt.Sprintf("%d", digit))
				//encList = append(encList, enc[0:len(enc)-5])
				fmt.Printf("\"%s\",\"%.6f\"\n",
					t.Text,
					float64(delta)/86400.0)
				/*
					fmt.Printf("%d  %s   %35s    %.6f _%d_\n",
						delta,
						enc[0:len(enc)-5],
						t.Text,
						float64(delta)/86400.0,
						digit)
						fmt.Printf("%d  %s   %35s    %.6f _%d_\n",
							delta,
							fmt.Sprintf("%d-%d", bs[0], bs[1]),
							t.Text,
							float64(delta)/86400.0,
							digit)*/
			}
			prev = t.Val
			//prevDigit = digit
		}

		/*
			for _, e := range encList {
				if e == "3" {
					fmt.Printf("%s ", ".")
				} else if e == "6" {
					fmt.Printf("%s ", "*")
				} else {
					fmt.Printf("%s ", " ")
				}
			}
			fmt.Println("")
		*/
		fmt.Println(len(all))
	} else if command == "next" {
		all := GetAll()
		fmt.Println(len(all))
	} else if command == "prev" {
		all := GetAll()
		fmt.Println(len(all))
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

func ParseData(f string) ([]Month, []int64) {
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
	deltas := []int64{}
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
			deltas = []int64{}
			times = []time.Time{}
		}
		times = append(times, u)
		if prev > 0 {
			delta = t.Val - prev
			deltas = append(deltas, delta)
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
	deltas = append(deltas, delta)
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
