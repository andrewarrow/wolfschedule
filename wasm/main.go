package main

import (
	"embed"
	"fmt"
	"io/fs"
	"syscall/js"
	"time"
)

//go:embed *.csv
var embededFiles embed.FS

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		//inputJSON := args[0].String()
		//fmt.Printf("input %s\n", inputJSON)
		return "pretty"
	})
	return jsonFunc
}

func main() {
	js.Global().Set("wolfData", jsonWrapper())

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

	for {
		time.Sleep(time.Second * 1)
	}
}
