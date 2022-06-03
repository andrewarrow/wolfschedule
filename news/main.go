package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/andrewarrow/wolfschedule/redis"
	"golang.org/x/net/html"
)

func main() {
	list := handleItems("raw.html")
	common := map[string]bool{}
	common["and"] = true
	common["the"] = true
	common["has"] = true
	common["she"] = true
	common["not"] = true
	common["didn't"] = true
	common["that"] = true
	common["this"] = true
	common["it's"] = true
	common["for"] = true
	common["what"] = true
	common["was"] = true
	common["are"] = true
	common["how"] = true
	common["why"] = true
	common["want"] = true
	common["our"] = true
	common["had"] = true
	common["onto"] = true
	common["new"] = true
	common["end"] = true
	common["with"] = true
	common["its"] = true
	common["will"] = true
	common["been"] = true
	common["near"] = true
	common["from"] = true
	reg, _ := regexp.Compile("[^a-z]+")

	m := map[string]bool{}
	for item, _ := range list {
		tokens := strings.Split(item, " ")
		for _, t := range tokens {
			if !unicode.IsLetter(rune(t[0])) {
				continue
			}
			lower := strings.ToLower(t)
			lower = reg.ReplaceAllString(lower, "")
			if len(lower) < 3 {
				continue
			}
			if common[lower] {
				continue
			}
			m[lower] = true
		}
	}
	for k, _ := range m {
		fmt.Printf("%s ", k)
	}
	fmt.Println("")
}
func main2() {
	list := handleItems("/home/aa/phantomjs/bin/raw.html")
	for item, href := range list {
		redis.InsertItem(time.Now().Unix(), item, href)
	}
}

func handleItems(filename string) map[string]string {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	tkn := html.NewTokenizer(strings.NewReader(s))

	hOn := false
	aOn := false
	aHref := ""

	list := map[string]string{}

	for {

		tt := tkn.Next()
		switch {

		case tt == html.ErrorToken:
			return list

		case tt == html.StartTagToken:

			t := tkn.Token()
			if t.Data == "h3" { // || t.Data == "h4" {
				hOn = true
			}
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key != "href" {
						continue
					}
					aHref = a.Val
				}
				aOn = true
			}

		case tt == html.TextToken:

			t := tkn.Token()
			txt := strings.TrimSpace(t.Data)
			if txt == "" {
				continue
			}
			if hOn && aOn {
				hOn = false
				aOn = false
				list[txt] = aHref
				aHref = ""
			}

		}

	}

	return list

}
