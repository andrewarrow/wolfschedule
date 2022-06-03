package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/wolfschedule/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Serve(os.Args[1])
}
