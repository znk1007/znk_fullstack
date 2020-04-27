package usertools

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(timeLayout, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}
