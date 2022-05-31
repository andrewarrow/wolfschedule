package main

import (
	"math/rand"
	"os"
	"time"
	"wolfschedule/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Serve(os.Args[1])
}
