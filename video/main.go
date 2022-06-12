package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "process" {
		dir := os.Args[2]
		ProcessDirectory(dir)
	} else if command == "test" {
		DrawOneFrame(time.Now())
	} else if command == "test2" {
		dir := os.Args[2]
		name := os.Args[3]
		DrawOnFrame(time.Now(), 345, dir, name)
	}
}
