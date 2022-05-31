package redis

import (
	"context"
	"fmt"
)

var ctx = context.Background()

func InsertItem(item string, ts int64) {
	err := nc().Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
