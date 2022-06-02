package moon

import (
	"fmt"
	"testing"
	"time"
)

func TestFindNextEvent(t *testing.T) {
	ts := time.Now().Unix()
	e := FindNextEvent(ts)
	fmt.Println(e.String())
}
