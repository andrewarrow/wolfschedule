package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/andrewarrow/wolfschedule/redis"
	"golang.org/x/net/html"
)

func main2() {
	list := handleItems("/home/aa/phantomjs/bin/raw.html")
	for item, href := range list {
		redis.InsertItem(time.Now().Unix(), item, href)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	list := handleItems("raw.html")
	t := time.Now()
	t = t.Add(time.Hour * -24)
	i := 0
	for {
		for item, href := range list {
			redis.InsertItem(t.Unix(), fmt.Sprintf("%d %s", rand.Intn(9), item), href)
		}
		fmt.Println(i, "done")
		time.Sleep(1 * time.Millisecond)
		t = t.Add(time.Hour * 1)
		i++
		if i > 24 {
			break
		}
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
