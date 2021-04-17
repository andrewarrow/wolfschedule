package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"sort"
	"strings"
)

type DigitAndIndex struct {
	Digit byte
	Index int
}

var mapByOne = map[string][]DigitAndIndex{}

func MakeSides(year string) {
	mapByOne = map[string][]DigitAndIndex{}
	_, deltas := ParseData(year + ".txt")

	myimage := image.NewRGBA(image.Rect(0, 0, 2350, 1450))
	mygreen := color.RGBA{0, 0, 0, 255}

	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)
	for i, d := range deltas {
		days := float64(d.Val) / 86400
		digit := AsciiByteToBase9(fmt.Sprintf("%d", d.Val))
		oneDigit := fmt.Sprintf("%.1f", days)
		MakeImage(myimage, i, days, digit, int(d.Time.Month()), d.Time.Day())
		mapByOne[oneDigit] = append(mapByOne[oneDigit], DigitAndIndex{digit, i})
	}
	myfile, _ := os.Create(fmt.Sprintf(year + ".jpg"))
	jpeg.Encode(myfile, myimage, &jpeg.Options{jpeg.DefaultQuality})

	fmt.Println("16|")
	fmt.Println("  |" + allThe("15.8", "15.9")) //.8 .9
	fmt.Println("  |" + allThe("15.6", "15.7")) //.7 .6
	fmt.Println("  |" + allThe("15.5"))
	fmt.Println("  |" + allThe("15.3", "15.4")) //.1 .2
	fmt.Println("  |" + allThe("15.1", "15.2")) //.1 .2
	fmt.Println("15|" + allThe("15.0"))
	fmt.Println("  |" + allThe("14.8", "14.9")) //.8 .9
	fmt.Println("  |" + allThe("14.6", "14.7")) //.7 .6
	fmt.Println("  |" + allThe("14.5"))
	fmt.Println("  |" + allThe("14.3", "14.4")) //.1 .2
	fmt.Println("  |" + allThe("14.1", "14.2")) //.1 .2
	fmt.Println("14|" + allThe("14.0"))
	fmt.Println("  |" + allThe("13.8", "13.9")) //.8 .9
	fmt.Println("  |" + allThe("13.6", "13.7")) //.7 .6
	fmt.Println("  |" + allThe("13.5"))
	fmt.Println("  |" + allThe("13.3", "13.4")) //.1 .2
	fmt.Println("  |" + allThe("13.1", "13.2")) //.1 .2
	fmt.Println("13|")
	fmt.Printf("   ")
	for _, d := range deltas {
		ms := fmt.Sprintf("%d", int(d.Time.Month()))
		if len(ms) == 1 {
			fmt.Printf(" %s ", ms)
		} else {
			fmt.Printf("%s ", ms)
		}
	}
	fmt.Println("")
	fmt.Printf("   ")
	for _, d := range deltas {
		ms := fmt.Sprintf("%d", d.Time.Day())
		if len(ms) == 1 {
			fmt.Printf(" %s ", ms)
		} else {
			fmt.Printf("%s ", ms)
		}
	}
	fmt.Println("")
}

func allThe(s ...string) string {
	list := []DigitAndIndex{}
	for _, item := range s {
		list = append(list, mapByOne[item]...)
	}
	if len(list) == 0 {
		return ""
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Index < list[j].Index
	})
	max := list[len(list)-1]
	m := map[int]byte{}
	for _, thing := range list {
		m[thing.Index] = thing.Digit
	}
	buff := []string{}
	for i := 0; i < max.Index+1; i++ {
		if m[i] != 0 {
			buff = append(buff, fmt.Sprintf(" %d", m[i]))
		} else {
			buff = append(buff, "  ")
		}
	}
	return strings.Join(buff, " ")
}
