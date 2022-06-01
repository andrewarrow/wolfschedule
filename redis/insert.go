package redis

import (
	"fmt"
	"time"
)

func InsertItem(ts int64, title, href string) {

	t := time.Unix(ts, 0)

	bucket := BucketForHour(t)

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
	SetBucketBoolean(title, bucket)

	expireTime := time.Now().Add(time.Hour * 24 * 30 * 12 * 2)
	nc().ExpireAt(ctx, bucket, expireTime)
	nc().ExpireAt(ctx, title, expireTime)

}

func SetBucketBoolean(title, bucket string) {
	err := nc().HSet(ctx, title, bucket, 1).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
