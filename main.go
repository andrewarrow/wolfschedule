package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("2021.txt")
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		fmt.Println(line)
	}
}
