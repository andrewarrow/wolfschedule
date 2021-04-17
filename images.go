package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func MakeImage(myimage *image.RGBA, i int, val float64, digit byte, month, day int) {
	//fmt.Println(i, oneDigit, digit, month, day)
	ity := 100
	if val < 14.1 && val >= 14.0 {
		ity = 56
	} else if val < 14.2 && val >= 14.1 {
		ity = 58
	} else if val < 14.3 && val >= 14.2 {
		ity = 60
	} else if val < 14.4 && val >= 14.3 {
		ity = 62
	} else if val < 14.5 && val >= 14.4 {
		ity = 64
	} else if val < 14.6 && val >= 14.5 {
		ity = 66
	} else if val < 14.7 && val >= 14.6 {
		ity = 68
	} else if val < 14.8 && val >= 14.7 {
		ity = 70
	} else if val < 14.9 && val >= 14.8 {
		ity = 72
	} else if val < 15.0 && val >= 14.9 {
		ity = 74
	} else if val < 15.1 && val >= 15.0 {
		ity = 76
	} else if val < 15.2 && val >= 15.1 {
		ity = 78
	} else if val < 15.3 && val >= 15.2 {
		ity = 80
	} else if val < 15.4 && val >= 15.3 {
		ity = 82
	} else if val < 15.5 && val >= 15.4 {
		ity = 84
	} else if val < 15.6 && val >= 15.5 {
		ity = 86
	} else if val < 14.0 && val >= 13.9 {
		ity = 54
	} else if val < 13.9 && val >= 13.8 {
		ity = 52
	} else if val < 13.8 && val >= 13.9 {
		ity = 50
	}
	per := float64(ity) / 100.0

	ten := 10 + (i * 100)
	red_rect := image.Rect(ten, 10, ten+100, 1400)
	myred := color.RGBA{0, 0, uint8(255.0 * per), 255}
	if digit == 3 {
		myred = color.RGBA{0, uint8(255.0 * per), 0, 255}
	} else if digit == 6 {
		myred = color.RGBA{uint8(255.0 * per), 0, 0, 255}
	}

	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

}
func MakeOneImage(i int, per float64) {

	myimage := image.NewRGBA(image.Rect(0, 0, 1850, 1450)) // x1,y1,  x2,y2
	mygreen := color.RGBA{0, 0, 0, 255}                    //  R, G, B, Alpha

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	red_rect := image.Rect(60, 80, 520, 560) //  geometry of 2nd rectangle
	myred := color.RGBA{0, 0, uint8(255.0 * per), 255}

	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

	myfile, _ := os.Create(fmt.Sprintf("out%d.jpg", i))
	jpeg.Encode(myfile, myimage, &jpeg.Options{jpeg.DefaultQuality})
}
