package moon

import (
	"fmt"
	"math"
	"time"
)

type Event struct {
	Timestamp int64
	FullMoon  bool
	Prev      *Event
	Next      *Event
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
		return fmt.Sprintf("FULL %s, Prev: %s", tstr, e.Prev.String())
	}
	return fmt.Sprintf("NEW %s", tstr)
}

func FindNextEvent(t int64) *Event {
	for i, k := range timeList {
		if k.Timestamp > t {
			k.Prev = &timeList[i-1]
			if len(timeList) > i+1 {
				k.Next = &timeList[i+1]
			}
			return &k
		}
	}

	return nil
}

func (e *Event) Until(t int64) string {
	delta := float64(e.Timestamp - t)
	d := delta / 86400.0
	days := math.Floor(d)
	hours := 24 * (d - days)
	fullHours := math.Floor(hours)
	mins := (hours - fullHours) * 60
	return fmt.Sprintf("%d day(s), %d hour(s), %d min(s)", int(days), int(fullHours), int(mins))
}

var timeList = []Event{
	Event{1653910320, false, nil, nil},
	Event{1655207520, true, nil, nil},
	Event{1656471180, false, nil, nil},
	Event{1657737480, true, nil, nil},
	Event{1659030900, false, nil, nil},
}
