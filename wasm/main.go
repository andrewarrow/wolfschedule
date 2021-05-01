package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed *.csv
var embededFiles embed.FS

func main() {
	fsys, err := fs.Sub(embededFiles, ".")
	if err != nil {
		panic(err)
	}
	f, err := fsys.Open("1970_2100.csv")
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 138492)
	f.Read(buff)
	fmt.Println(string(buff))
}
