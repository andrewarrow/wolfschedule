package moon

import "fmt"

func FindNextEvent(t int64) {
	for k, v := range times {
		fmt.Println(k, v)
	}
}

var times map[int64]bool = map[int64]bool{
	1653910320: false,
	1655207520: true,
	1656471180: false,
	1657737480: true,
	1659030900: false,
}
