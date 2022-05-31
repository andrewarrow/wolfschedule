package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/wolfschedule/parse"
	"github.com/andrewarrow/wolfschedule/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	server.Year = time.Now().Year()
	timeZone, _ := time.LoadLocation("America/Phoenix")
	server.EventDate = time.Date(server.Year, time.Month(1), 1, 0, 0, 0, 0, timeZone)
	server.Month = server.EventDate.Month()

	all := parse.GetAll()
	for _, t := range all {
		u := fmt.Sprintf("%v", time.Unix(t.Val, 0))
		server.Times[u[0:10]] = t.Val
	}

	server.Serve(os.Args[1])
}
