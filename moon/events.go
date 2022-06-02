package moon

import (
	"fmt"
	"time"
)

type Event struct {
	Timestamp int64
	FullMoon  bool
}

func NewEvent(ts int64, b bool) *Event {
	e := Event{}
	e.Timestamp = ts
	e.FullMoon = b
	return &e
}

func (e *Event) String() string {
	t := time.Unix(e.Timestamp, 0)
	tstr := fmt.Sprintf("%v", t)
	if e.FullMoon {
		return fmt.Sprintf("FULL %s", tstr)
	}
	return fmt.Sprintf("NEW %s", tstr)
}

func FindNextEvent(t int64) *Event {
	for _, k := range timeList {
		if k.Timestamp > t {
			return &k
		}
	}

	return nil
}

var timeList = []Event{
	Event{1653910320, false},
	Event{1655207520, true},
	Event{1656471180, false},
	Event{1657737480, true},
	Event{1659030900, false},
}
