package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("html.txt")
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, " UTC") {
			tokens := strings.Split(line, "<p>")
			//for _, t := range tokens {
			//	fmt.Println(t)
			//}
			fmt.Println(tokens[1])
			fmt.Println(tokens[2])
		}
	}
	fmt.Println("vim-go", len(b))
}
