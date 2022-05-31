package redis

import (
	"fmt"
	"time"
)

func QueryDay() []string {
	t := time.Now()
	t = t.Add(time.Hour * -24)
	buckets := []string{}
	i := 0
	for {
		bucket := bucketForHour(t)
		buckets = append(buckets, bucket)
		t = t.Add(time.Hour)
		i++
		if i >= 24 {
			break
		}
	}

	items := []string{}
	for _, b := range buckets {
		members, err := nc().SMembers(ctx, b).Result()
		if err != nil {
			fmt.Println(err)
			return items
		}
		for _, i := range members {
			items = append(items, i)
		}
	}

	return items
}
