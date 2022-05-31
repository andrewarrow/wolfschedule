package redis

import (
	"fmt"
	"time"
)

func InsertItem(ts int64, title, href string) {

	t := time.Unix(ts, 0)

	bucket := bucketForHour(t)

	err := nc().SAdd(ctx, bucket, title).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = nc().HSet(ctx, title, "href", href).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
	nc().ExpireAt(ctx, title, expireTime)

}
