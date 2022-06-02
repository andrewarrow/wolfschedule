package moon

import (
	"testing"
	"time"
)

func TestFindNextEvent(t *testing.T) {
	ts := time.Now().Unix()
	FindNextEvent(ts)
}
