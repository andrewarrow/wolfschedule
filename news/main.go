package main

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/redis"
	"golang.org/x/net/html"
)

func main() {
	list := handleItems("/home/aa/phantomjs/bin/raw.html")
	for item, _ := range list {
		redis.InsertItem(time.Now().Unix(), item)
	}

}

func handleItems(filename string) map[string]bool {
	b, _ := ioutil.ReadFile(filename)
	s := string(b)
	tkn := html.NewTokenizer(strings.NewReader(s))

	hOn := false
	aOn := false

	list := map[string]bool{}

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
				list[txt] = true
			}

		}

	}

	return list

}
