package main

import (
	"fmt"
	"time"

	"github.com/fogleman/gg"
)

func DrawOneFrame(modtime time.Time) {
	dc := gg.NewContext(1020, 576)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	logo, e := gg.LoadPNG("logo.png")
	fmt.Println(e)

	pattern := gg.NewSurfacePattern(logo, gg.RepeatNone)
	dc.MoveTo(0, 0)
	dc.LineTo(1020, 0)
	dc.LineTo(1020, 576)
	dc.LineTo(0, 576)
	dc.LineTo(0, 0)
	dc.SetFillStyle(pattern)
	dc.Fill()

	dc.LoadFontFace("arial.ttf", 24)

	i := 1
	t := modtime.Add(time.Second * time.Duration(i))
	dc.DrawStringAnchored(t.Format(time.RFC850), 310, 520, 0.5, 0.5)

	dc.SavePNG("test.png")
}
