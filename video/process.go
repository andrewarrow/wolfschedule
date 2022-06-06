package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func ProcessDirectory(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		name := file.Name()
		fmt.Println(name)
		tokens := strings.Split(name, "_")
		if tokens[0] != "IMG" {
			continue
		}
		parts := strings.Split(tokens[1], ".")
		outputDir := fmt.Sprintf("%s/DONE_%s", dir, parts[0])
		os.Mkdir(outputDir, 0755)
		outputFilename := fmt.Sprintf("DONE_%s/source.mov", parts[0])
		details := "scale=w=880:h=720:force_original_aspect_ratio=1,pad=880:720:(ow-iw):(oh-ih)"
		cmd := exec.Command("/usr/local/bin/ffmpeg", "-i", name, "-vf", details, outputFilename)
		cmd.Dir = dir
		output, e := cmd.CombinedOutput()
		if e != nil {
			fmt.Println(name, e)
			return
		}
		fmt.Println(name, string(output))
	}
}

//  ffmpeg -i IMG_2367.MOV -vf "scale=w=880:h=720:force_original_aspect_ratio=1,pad=880:720:(ow-iw):(oh-ih)" output.mov
// ffmpeg -i DONE_2367.MOV -vf fps=29.97 img%05d.png
//  ffmpeg -framerate 29.97 -pattern_type glob -i '*.png' -c:v libx264 -pix_fmt yuv420p out.mp4
// ffmpeg -i out.mp4 -i DONE_2367.mp3 -c copy -map 0:v:0 -map 1:a:0 sound.mov
