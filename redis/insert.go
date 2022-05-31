package redis

import (
	"fmt"
)

func InsertItem(ts int64, item string) {

	err := nc().SAdd(ctx, item, fmt.Sprintf("%d", ts)).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
