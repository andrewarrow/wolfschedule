package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"time"
)

func DisplayCurrentDay(year string, add int) int64 {
	//y, _ := strconv.Atoi(year)
	now := time.Now()
	if year == "" {
		//	y = now.Year()
	}
	now = now.AddDate(0, 0, add)
	today := fmt.Sprintf("%v", now)
	all := GetAll()
	m := map[string]int64{}
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = t.Val
	}
	day1 := now.AddDate(0, 0, -25)
	b1 := day1.AddDate(0, 0, +38)
	last := int64(0)
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
	return last
}

func MakePDF(year string, month int) {
	all := GetAll()
	m := map[string]string{}
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		m[u[0:10]] = "0"
	}
	day0, _ := time.Parse(special, fmt.Sprintf("01/01/2003 00:00"))
	day1 := day0
	day365, _ := time.Parse(special, fmt.Sprintf("12/31/2030 00:00"))
	eventHappened := -1
	for {
		u := fmt.Sprintf("%v", day1)
		substring := u[0:10]
		if m[substring] == "0" {
			eventHappened = 0
			inner := day1
			inner = inner.AddDate(0, 0, -5)
			innerU := fmt.Sprintf("%v", inner)
			innerSub := innerU[0:10]
			m[innerSub] = "."
		} else if eventHappened == 5 {
			m[substring] = "."
		}
		eventHappened++
		day1 = day1.AddDate(0, 0, 1)
		if day1.After(day365) {
			break
		}
	}

	eventHappened = 0
	buff := []string{}
	day1, _ = time.Parse(special, fmt.Sprintf("%02d/01/%s 00:00", month, year))
	day1 = day1.AddDate(0, 0, -5)

	myimage := image.NewRGBA(image.Rect(0, 0, 1000, 1500))
	mygreen := color.RGBA{255, 255, 255, 255}
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)
	for {
		u := fmt.Sprintf("%v", day1)
		wd := fmt.Sprintf("%v", day1.Weekday())
		if wd == "Tuesday" || wd == "Thursday" || wd == "Saturday" ||
			wd == "Sunday" {
			wd = ""
		}

		col1 := "" // event date
		col2 := "" // normal date
		col4 := "" // wd

		substring := u[0:10]
		if m[substring] == "0" {
			col1 = substring
			eventHappened++
		} else {
			col2 = substring
		}
		col4 = wd
		thing := fmt.Sprintf("%10s %10s%20s%30s", col1, col2, m[substring], col4)
		buff = append(buff, thing)
		fmt.Println(thing)

		red_rect := image.Rect(60, 80, 520, 560)
		myred := color.RGBA{0, 0, 0, 255}
		draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

		if eventHappened == 3 && m[substring] == "." {
			break
		}

		day1 = day1.AddDate(0, 0, 1)
	}

	myfile, _ := os.Create(fmt.Sprintf("html/%d.jpg", month))
	jpeg.Encode(myfile, myimage, &jpeg.Options{jpeg.DefaultQuality})

}
