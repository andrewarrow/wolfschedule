package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	ParseData("2021.txt")
	ParseData("2022.txt")
}

func ParseData(f string) {
	b, _ := ioutil.ReadFile(f)
	s := string(b)
	for _, line := range strings.Split(s, "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			break
		}
		newMoon := tokens[0] + " " + tokens[1]
		fullMoon := tokens[2] + " " + tokens[3]

		newDate, _ := time.Parse("01/02/2006 15:04", newMoon)
		fullDate, _ := time.Parse("01/02/2006 15:04", fullMoon)

		fmt.Println(newDate, fullDate)
		delta := fullDate.Unix() - newDate.Unix()
		deltaString := fmt.Sprintf("%d", delta)
		fmt.Println(newDate.Unix(), fullDate.Unix(), delta,
			AsciiByteToBase9(deltaString))

	}
}
func AsciiByteToBase9(a string) byte {

	sum := byte(0)
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	if len(strSum) > 1 {
		return AsciiByteToBase9(strSum)
	}
	return sum

}
