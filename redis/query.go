package redis

import (
	"fmt"
	"time"
)

func QueryDay() []string {
	t := time.Now()
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

	buckets = []string{"2022053018",
		"2022053020",
		"2022053016",
		"2022053019",
		"2022053021",
		"2022053017",
		"2022053022",
		"2022053015"}

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
