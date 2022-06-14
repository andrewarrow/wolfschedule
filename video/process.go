package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/moon"
	"github.com/fogleman/gg"
)

func ProcessDirectory(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		name := file.Name()
		modifiedtime := file.ModTime()
		fmt.Println(name, modifiedtime)
		tokens := strings.Split(name, "_")
		if tokens[0] != "IMG" {
			continue
		}
		parts := strings.Split(tokens[1], ".")
		ScaleOrig(dir, name, parts[0])
		ExtractFrames(dir, parts[0], modifiedtime)
		ExtractAudio(dir, parts[0])
		AssembleFromFrames(dir, parts[0])
		AddBackSound(dir, parts[0])
	}
}

func ScaleOrig(dir, name, part string) {
	outputDir := fmt.Sprintf("%s/DONE_%s", dir, part)
	os.Mkdir(outputDir, 0755)
	outputFilename := fmt.Sprintf("DONE_%s/source.mov", part)
	details := "scale=w=1020:h=576:force_original_aspect_ratio=1,pad=1020:576:(ow-iw):(oh-ih)"
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", name, "-vf", details, outputFilename)
	cmd.Dir = dir
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(name, e)
		return
	}
	fmt.Println(name, string(output))
}

func ExtractFrames(dir, part string, modtime time.Time) {
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", "source.mov", "-vf", "fps=29.97", "img%07d.png")
	cmd.Dir = fmt.Sprintf("%s/DONE_%s", dir, part)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(part, e)
		return
	}
	fmt.Println(part, string(output))
	files, _ := ioutil.ReadDir(cmd.Dir)
	sum := 0
	for i, file := range files {
		name := file.Name()
		if !strings.HasPrefix(name, "img0") {
			continue
		}
		fmt.Println(name)
		DrawOnFrame(modtime, sum, cmd.Dir, name)
		if i%30 == 0 {
			sum++
		}
	}
}

func DrawOnFrame(modtime time.Time, i int, dir, name string) {
	dc := gg.NewContext(1020, 576)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	path := fmt.Sprintf("%s/%s", dir, name)
	fmt.Println(path)
	existing, e := gg.LoadPNG(path)
	fmt.Println(e)
	logo, e := gg.LoadPNG("logo.png")
	fmt.Println(e)

	pattern := gg.NewSurfacePattern(existing, gg.RepeatNone)
	dc.MoveTo(0, 0)
	dc.LineTo(1020, 0)
	dc.LineTo(1020, 576)
	dc.LineTo(0, 576)
	dc.LineTo(0, 0)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	pattern = gg.NewSurfacePattern(logo, gg.RepeatNone)
	dc.MoveTo(0, 0)
	dc.LineTo(1020, 0)
	dc.LineTo(1020, 576)
	dc.LineTo(0, 576)
	dc.LineTo(0, 0)
	dc.SetFillStyle(pattern)
	dc.Fill()

	dc.LoadFontFace("arial.ttf", 24)

	t := modtime.Add(time.Second * time.Duration(i))
	event := moon.FindNextEvent(t.Unix())
	str := fmt.Sprintf("Next %s MOON in", event.NewOrFull())
	dc.DrawStringAnchored(str, 310, 430, 0.5, 0.5)
	deltaString := moon.EventDelta(event.Timestamp - t.Unix())
	dc.DrawStringAnchored(deltaString, 310, 460, 0.5, 0.5)
	dc.DrawStringAnchored(t.Format(time.RFC850), 310, 520, 0.5, 0.5)

	dc.SavePNG(path)
}

func ExtractAudio(dir, part string) {
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", "source.mov", "audio.mp3")
	cmd.Dir = fmt.Sprintf("%s/DONE_%s", dir, part)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(part, e)
		return
	}
	fmt.Println(part, string(output))
}

func AssembleFromFrames(dir, part string) {
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-framerate", "29.97", "-pattern_type", "glob", "-i", "*.png", "-c:v", "libx264",
		"-pix_fmt", "yuv420p", "temp.mov")
	cmd.Dir = fmt.Sprintf("%s/DONE_%s", dir, part)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(part, e)
		return
	}
	fmt.Println(part, string(output))
}

func AddBackSound(dir, part string) {
	cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", "temp.mov", "-i", "audio.mp3", "-c", "copy", "-map", "0:v:0", "-map", "1:a:0", "sound.mov")
	cmd.Dir = fmt.Sprintf("%s/DONE_%s", dir, part)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Println(part, e)
		return
	}
	fmt.Println(part, string(output))
}

//  ffmpeg -i IMG_2367.MOV -vf "scale=w=880:h=720:force_original_aspect_ratio=1,pad=880:720:(ow-iw):(oh-ih)" output.mov
// ffmpeg -i DONE_2367.MOV -vf fps=29.97 img%05d.png
//  ffmpeg -framerate 29.97 -pattern_type glob -i '*.png' -c:v libx264 -pix_fmt yuv420p out.mp4
// ffmpeg -i out.mp4 -i DONE_2367.mp3 -c copy -map 0:v:0 -map 1:a:0 sound.mov

//file '/Users/aa/Desktop/meditation/DONE_2367/sound.mov'
//file '/Users/aa/wolfschedule/beeps/b2.mov'
//file '/Users/aa/Desktop/meditation/DONE_2368/sound.mov'
//  ffmpeg -f concat -safe 0 -i mylist.txt -c copy output.mov
