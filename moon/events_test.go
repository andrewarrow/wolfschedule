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
	list := FindEventsForYear(2022)
	for _, item := range list {
		fmt.Println(item.Timestamp)
	}
}
