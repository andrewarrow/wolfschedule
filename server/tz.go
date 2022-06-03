package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TimeZonePost(c *gin.Context) {
	tz := c.PostForm("tz")
	c.SetCookie("tz", tz, 3600*24*365*10, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

var zones = []string{"America/Anchorage",
	"America/Chicago",
	"America/Denver",
	"America/Detroit",
	"America/Jamaica",
	"America/Los_Angeles",
	"America/New_York",
	"America/Panama",
	"America/Phoenix",
	"Antarctica/Troll",
	"Asia/Bangkok",
	"Asia/Dubai",
	"Asia/Hong_Kong",
	"Asia/Kabul",
	"Asia/Kuwait",
	"Asia/Pyongyang",
	"Asia/Qatar",
	"Asia/Seoul",
	"Asia/Singapore",
	"Asia/Tehran",
	"Asia/Tokyo",
	"Australia/Brisbane",
	"Australia/Melbourne",
	"Australia/Sydney",
	"Europe/Amsterdam",
	"Europe/Athens",
	"Europe/Belgrade",
	"Europe/Brussels",
	"Europe/Budapest",
	"Europe/Copenhagen",
	"Europe/Dublin",
	"Europe/Helsinki",
	"Europe/London",
	"Europe/Minsk",
	"Europe/Monaco",
	"Europe/Paris",
	"Europe/Sarajevo",
	"Europe/Stockholm",
	"Europe/Vatican",
	"Europe/Warsaw",
	"Europe/Zurich",
	"Indian/Cocos",
	"Pacific/Fiji",
	"Pacific/Guam",
	"Pacific/Honolulu",
	"UTC"}
