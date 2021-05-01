package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

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
	day1Orig := day1
	day1 = day1.AddDate(0, 0, -5)

	//2550 x 3300
	myimage := image.NewRGBA(image.Rect(0, 0, 1275, 1650))
	mygreen := color.RGBA{255, 255, 255, 255}
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)
	addLabel(myimage, 650, 80, fmt.Sprintf("%v %d", day1Orig.Month(), day1Orig.Year()))
	row := 0
	offset := 100
	rowSize := 40
	leftPush := 100
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
			red_rect := image.Rect(50+leftPush, 10+((row-1)*rowSize)+offset+1, 159+leftPush, 10+rowSize+((row-1)*rowSize)+offset+1)
			myred := color.RGBA{222, 128, 222, 255}
			draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)
		} else {
			col2 = substring
		}
		col4 = wd
		thing := fmt.Sprintf("%10s %10s%20s%30s", col1, col2, m[substring], col4)
		buff = append(buff, thing)
		fmt.Println(thing)

		red_rect := image.Rect(50+leftPush, 10+(row*rowSize)+offset, 1200, 11+(row*rowSize)+offset)
		myred := color.RGBA{0, 0, 0, 255}
		draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)
		addLabel(myimage, 1100, 10+(row*rowSize)+offset-10, col4)
		if col2 != "" {
			addLabel(myimage, 160+leftPush, 10+(row*rowSize)+offset-10+4, col2)
		} else if col1 != "" {
			addLabel(myimage, 60+leftPush, 10+(row*rowSize)+offset-10+4, col1)
		}

		if row >= 32 && int(day1Orig.Month()) != int(day1.Month()) {
			break
		}

		day1 = day1.AddDate(0, 0, 1)
		row++
	}

	myfile, _ := os.Create(fmt.Sprintf("html/%s_%d.jpg", year, month))
	jpeg.Encode(myfile, myimage, &jpeg.Options{100})

}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{25, 25, 25, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Bold8x16,
		Dot:  point,
	}
	d.DrawString(label)
}
