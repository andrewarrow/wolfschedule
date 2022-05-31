package redis

import (
	"context"
	"fmt"
)

var ctx = context.Background()

func InsertItem(ts int64, item string) {

	err := nc().Set(ctx, item, fmt.Sprintf("%d", ts), 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
