package redis

import (
	"fmt"
	"sort"
	"time"
)

type Item struct {
	Count int
	Title string
}

func QueryDay() []Item {
	t := time.Now()
	t = t.Add(time.Hour)
	t = t.Add(time.Hour * -24)
	buckets := []string{}
	i := 0
	for {
		bucket := bucketForHour(t)
		buckets = append([]string{bucket}, buckets...)
		t = t.Add(time.Hour)
		i++
		if i >= 24 {
			break
		}
	}

	m := map[string]map[string]int{}
	for _, b := range buckets {
		m[b] = map[string]int{}
		for _, item := range QueryBucket(b) {
			m[b][item] = 1
		}
	}

	freshest := buckets[0]
	i = 0
	for {
		if len(m[freshest]) > 0 || i >= len(buckets)-1 {
			break
		}
		i++
		freshest = buckets[i]
	}
	rest := buckets[i+1:]

	for k, _ := range m[freshest] {
		for _, history := range rest {
			if m[history][k] == 1 {
				m[freshest][k]++
			}
		}
	}

	items := []Item{}
	for k, v := range m[freshest] {
		item := Item{v, k}
		items = append(items, item)
	}
	sort.Slice(items, func(a, b int) bool {
		return items[a].Count > items[b].Count
	})
	return items
}

func QueryBucket(b string) []string {
	items := []string{}
	members, err := nc().SMembers(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return items
	}
	for _, i := range members {
		items = append(items, i)
	}
	return items
}

func QueryAttributes(b string) map[string]string {
	m, err := nc().HGetAll(ctx, b).Result()
	if err != nil {
		fmt.Println(err)
		return m
	}
	return m
}
