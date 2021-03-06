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

func (e *Event) Clone() *Event {
	clone := NewEvent(e.Timestamp, e.FullMoon)
	return clone
}

func (e *Event) AsTime(location *time.Location) time.Time {
	t := time.Unix(e.Timestamp, 0)
	return t.In(location)
}

func (e *Event) String() string {
	t := time.Unix(e.Timestamp, 0)
	tstr := fmt.Sprintf("%v", t)
	if e.FullMoon {
		return fmt.Sprintf("FULL %s, Prev: %s", tstr, e.Prev.String())
	}
	return fmt.Sprintf("NEW %s", tstr)
}

func (e *Event) NewOrFull() string {
	if e.FullMoon {
		return "FULL"
	}
	return "NEW"
}

func FindNextEvent(t int64) *Event {
	for i, k := range timeList {
		if k.Timestamp > t {
			e := k.Clone()
			if i > 0 {
				e.Prev = timeList[i-1].Clone()
			}
			if len(timeList) > i+1 {
				e.Next = timeList[i+1].Clone()
			}
			return e
		}
	}

	return nil
}

type ChartData struct {
	Deltas []int
	Labels []string
}

func FindEventsForChart() ChartData {
	cd := ChartData{}
	for i := len(timeList) - 1; i > 1; i-- {
		delta := int(timeList[i].Timestamp - timeList[i-1].Timestamp)
		cd.Deltas = append([]int{delta / 60}, cd.Deltas...)
		t := time.Unix(timeList[i].Timestamp, 0)
		u := fmt.Sprintf("%v", t)
		substring := u[0:10]
		cd.Labels = append([]string{substring}, cd.Labels...)
	}

	return cd
}

func FindEventsForYear(year int, location *time.Location) map[string]*Event {
	m := map[string]*Event{}
	on := false
	for _, k := range timeList {
		t := time.Unix(k.Timestamp, 0)
		t = t.In(location)
		if t.Year() == year {
			on = true
			u := fmt.Sprintf("%v", t)
			substring := u[0:10]
			m[substring] = k.Clone()
		} else if on {
			return m
		}
	}

	return m
}

func EventDelta(t int64) string {
	delta := float64(t)
	d := delta / 86400.0
	days := math.Floor(d)
	hours := 24 * (d - days)
	fullHours := math.Floor(hours)
	mins := (hours - fullHours) * 60
	return fmt.Sprintf("%d day(s), %d hour(s), %d min(s)", int(days), int(fullHours), int(mins))
}

var timeList = []Event{
	Event{1610514120, false, nil, nil},
	Event{1611861480, true, nil, nil},
	Event{1613070480, false, nil, nil},
	Event{1614413940, true, nil, nil},
	Event{1615630980, false, nil, nil},
	Event{1616957400, true, nil, nil},
	Event{1618194720, false, nil, nil},
	Event{1619494380, true, nil, nil},
	Event{1620759660, false, nil, nil},
	Event{1622027640, true, nil, nil},
	Event{1623322440, false, nil, nil},
	Event{1624560000, true, nil, nil},
	Event{1625879820, false, nil, nil},
	Event{1627094220, true, nil, nil},
	Event{1628430600, false, nil, nil},
	Event{1629633720, true, nil, nil},
	Event{1630975920, false, nil, nil},
	Event{1632182040, true, nil, nil},
	Event{1633518300, false, nil, nil},
	Event{1634741820, true, nil, nil},
	Event{1636060500, false, nil, nil},
	Event{1637312340, true, nil, nil},
	Event{1638603840, false, nil, nil},
	Event{1639888620, true, nil, nil},
	Event{1641148500, false, nil, nil},
	Event{1642463460, true, nil, nil},
	Event{1643694540, false, nil, nil},
	Event{1645030740, true, nil, nil},
	Event{1646242680, false, nil, nil},
	Event{1647588000, true, nil, nil},
	Event{1648794420, false, nil, nil},
	Event{1650135420, true, nil, nil},
	Event{1651350600, false, nil, nil},
	Event{1652674500, true, nil, nil},
	Event{1653910320, false, nil, nil},
	Event{1655207520, true, nil, nil},
	Event{1656471180, false, nil, nil},
	Event{1657737480, true, nil, nil},
	Event{1659030900, false, nil, nil},
	Event{1660268160, true, nil, nil},
	Event{1661588160, false, nil, nil},
	Event{1662803880, true, nil, nil},
	Event{1664142840, false, nil, nil},
	Event{1665348840, true, nil, nil},
	Event{1666694880, false, nil, nil},
	Event{1667905320, true, nil, nil},
	Event{1669244220, false, nil, nil},
	Event{1670472540, true, nil, nil},
	Event{1671790620, false, nil, nil},
	Event{1673046540, true, nil, nil},
	Event{1674334500, false, nil, nil},
	Event{1675621800, true, nil, nil},
	Event{1676876940, false, nil, nil},
	Event{1678192920, true, nil, nil},
	Event{1679419560, false, nil, nil},
	Event{1680755820, true, nil, nil},
	Event{1681964100, false, nil, nil},
	Event{1683308160, true, nil, nil},
	Event{1684511700, false, nil, nil},
	Event{1685850180, true, nil, nil},
	Event{1687063140, false, nil, nil},
	Event{1688384400, true, nil, nil},
	Event{1689618780, false, nil, nil},
	Event{1690914780, true, nil, nil},
	Event{1692178680, false, nil, nil},
	Event{1693445820, true, nil, nil},
	Event{1694742000, false, nil, nil},
	Event{1695981480, true, nil, nil},
	Event{1697306100, false, nil, nil},
	Event{1698524640, true, nil, nil},
	Event{1699867620, false, nil, nil},
	Event{1701076560, true, nil, nil},
	Event{1702423920, false, nil, nil},
	Event{1703637180, true, nil, nil},
}
