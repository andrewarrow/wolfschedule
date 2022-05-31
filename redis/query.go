package redis

import (
	"fmt"
	"time"
)

func QueryDay() {
	t := time.Now()
	buckets := []string{}
	i := 0
	for {
		bucket := bucketForHour(t)
		buckets = append(buckets, bucket)
		t = t.Add(time.Hour)
		i++
		if i > 24 {
			break
		}
	}

	fmt.Println(buckets)
}
