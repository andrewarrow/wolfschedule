package redis

import (
	"fmt"
	"time"
)

func InsertItem(ts int64, item string) {

	t := time.Unix(ts, 0)

	bucket := bucketForHour(t)

	err := nc().SAdd(ctx, bucket, item).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
