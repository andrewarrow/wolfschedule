package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeAndZone(c *gin.Context) (int64, string) {
	t := c.DefaultQuery("t", fmt.Sprintf("%d", time.Now().Unix()))
	tInt, _ := strconv.ParseInt(t, 10, 64)

	tz, _ := c.Cookie("tz")
	if tz == "" {
		tz = "Antarctica/Troll"
	}

	return tInt, tz
}
