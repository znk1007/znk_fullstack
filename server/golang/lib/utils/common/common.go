package utils

import (
	"strconv"
	"time"
)

// GetNowTime 获取当前时间戳
func GetNowTime() string {
	unixTime := time.Now().Unix()
	return strconv.FormatInt(unixTime, 10)
}

// DeleteString 删除指定元素
func DeleteString(slices []string, ele string) []string {
	i := 0
	for _, val := range slices {
		if val != ele {
			slices[i] = val
			i++
		}
	}
	return slices[:i]
}
