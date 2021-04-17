package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func MakeImages() {

	myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2
	mygreen := color.RGBA{0, 100, 0, 255}                //  R, G, B, Alpha

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	red_rect := image.Rect(60, 80, 120, 160) //  geometry of 2nd rectangle
	myred := color.RGBA{200, 0, 0, 255}

	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

	myfile, _ := os.Create("out.jpg")
	jpeg.Encode(myfile, myimage, &jpeg.Options{jpeg.DefaultQuality})
}
