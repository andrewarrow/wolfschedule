package redis

import (
	"fmt"
	"testing"
)

func TestQueryDay(t *testing.T) {
	items := QueryDay()
	for _, item := range items {
		fmt.Println(item)
	}
}
