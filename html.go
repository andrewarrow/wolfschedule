package main

import "fmt"

func MakeHtml(months []Month) {
	fmt.Println("<pre>")
	for _, m := range months {
		fmt.Println(m.HTML() + "<br/>")
	}
	fmt.Println("</pre>")
}
