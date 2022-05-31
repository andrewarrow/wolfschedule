package redis

import (
	"fmt"
	"time"
)

func InsertItem(ts int64, item string) {

	t := time.Unix(ts, 0)

	bucketForHour := bucket_for_hour(t)
	//bucket_for_day := bucket_for_day(t)
	//bucket_for_month := bucket_for_month(t)
	//bucket_for_year := bucket_for_year(t)

	err := nc().SAdd(ctx, bucketForHour, item).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}

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
func bucket_for_hour(t time.Time) string {
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
