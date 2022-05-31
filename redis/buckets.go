package redis

import (
	"fmt"
	"time"
)

func bucket_for_year(t time.Time) string {
	format := t.Format("2006")
	return fmt.Sprintf("%s", format)
}
func bucket_for_month(t time.Time) string {
	format := t.Format("200601")
	return fmt.Sprintf("%s", format)
}
func bucket_for_day(t time.Time) string {
	format := t.Format("20060102")
	return fmt.Sprintf("%s", format)
}
func bucketForHour(t time.Time) string {
	format := t.Format("2006010215")
	return fmt.Sprintf("%s", format)
}
func bucket_for_min(t time.Time) string {
	format := t.Format("200601021504")
	return fmt.Sprintf("%s", format)
}
func bucket_for_sec(t time.Time) string {
	format := t.Format("20060102150405")
	return fmt.Sprintf("%s", format)
}
