package parse

import (
	"fmt"
	"io/fs"
	"sort"
	"strconv"
	"strings"
	"time"
)

var special = "01/02/2006 15:04"

type Thing struct {
	Text string
	Val  int64
}

var things = []Thing{}

func ForHTML() string {
	return things[0].Text
}

func LoadCSV(filename string, fsys fs.FS) {

	f, err := fsys.Open("1970_2100.csv")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 138492)
	f.Read(b)

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
func GetAll(fsys fs.FS) []Thing {
	things = []Thing{}
	LoadCSV("1970_2100.csv", fsys)
	sort.SliceStable(things, func(i, j int) bool {
		return things[i].Val < things[j].Val
	})
	return things
}

func DisplayCurrentDay(year string, add int) (int64, int64) {
	//y, _ := strconv.Atoi(year)
	now := time.Now()
	if year == "" {
		//	y = now.Year()
	}
	now = now.AddDate(0, 0, add)
	today := fmt.Sprintf("%v", now)
	m := map[string]int64{}
	for _, t := range things {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = t.Val
	}
	day1 := now.AddDate(0, 0, -25)
	b1 := day1.AddDate(0, 0, +40)
	last := int64(0)
	other := int64(0)
	for {
		u := fmt.Sprintf("%v", day1)
		wd := fmt.Sprintf("%v", day1.Weekday())
		if wd == "Tuesday" || wd == "Thursday" || wd == "Saturday" ||
			wd == "Sunday" {
			wd = ""
		}

		col1 := "" // event date
		col2 := "" // normal date
		col3 := "" // arrow
		col4 := "" // wd

		substring := u[0:10]
		if m[substring] > 0 {
			col1 = substring
			other = last
			last = m[substring]
		} else {
			col2 = substring
		}
		if substring == today[0:10] {
			col3 = " <-------------"
		}
		col4 = wd
		fmt.Printf("%10s %10s%20s%30s\n", col1, col2, col3, col4)
		day1 = day1.AddDate(0, 0, 1)
		if day1.After(b1) {
			break
		}
	}
	return last, other
}
