package main

import (
	"fmt"
	"os"
)

func main() {
	api := os.Getenv("CMC")
	fmt.Println("key", api)
}
