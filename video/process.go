package main

import (
	"fmt"
	"io/ioutil"
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
		details := "scale=w=880:h=720:force_original_aspect_ratio=1,pad=880:720:(ow-iw):(oh-ih)"
		outputFilename := fmt.Sprintf("DONE_%s", tokens[1])
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
