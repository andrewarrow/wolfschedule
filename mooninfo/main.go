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
			m := tokens[1]
			t := tokens[2]
			if len(m) < 50 {
				tokens = strings.Split(m, "<")
				md := tokens[0]
				tokens = strings.Split(t, "<")
				fmt.Println(md, tokens[0])
			}
		}
	}
}
