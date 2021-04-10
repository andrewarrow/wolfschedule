package main

import "fmt"

func MakeHtml(months []Month) {
	for _, m := range months {
		fmt.Println(m.String())
	}
}
