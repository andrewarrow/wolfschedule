package main

import (
	"embed"
	"io/fs"
	"syscall/js"
	"time"
	"wasm/parse"
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
		return parse.ForHTML()
	})
	return jsonFunc
}
func handleKey() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		parse.HandleKey(args[0].String())
		return ""
	})
	return jsonFunc
}

func main() {
	js.Global().Set("wolfData", jsonWrapper())
	js.Global().Set("handleKey", handleKey())

	fsys, err := fs.Sub(embededFiles, ".")
	if err != nil {
		panic(err)
	}

	parse.GetAll(fsys)

	for {
		time.Sleep(time.Second * 1)
	}
}
