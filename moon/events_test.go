package moon

import (
	"fmt"
	"testing"
	"time"
)

func TestFindNextEvent(t *testing.T) {
	ts := time.Now().Unix()
	e := FindNextEvent(ts)
	fmt.Println(e.String())
}

func TestFindEventsForYear(t *testing.T) {
	location, _ := time.LoadLocation("UTC")
	list := FindEventsForYear(2022, location)
	for k, item := range list {
		fmt.Println(k, item.Timestamp)
	}
}
