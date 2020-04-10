package usermiddleware

import (
	"fmt"
	"time"
)

//FreqAccess 访问频度
type FreqAccess struct {
	maxCount    int            //最大统计次数
	maxDuration int            //最大时间限制
	start       int64          //开始时刻
	access      map[string]int //当前统计次数
}

//NewFreqAccess 初始化访问频度对象
func NewFreqAccess(maxDuration int, maxCount int) *FreqAccess {
	return &FreqAccess{
		maxDuration: maxDuration,
		maxCount:    maxCount,
		access:      make(map[string]int),
	}
}

//AccessCtrl 访问控制
func (fa *FreqAccess) AccessCtrl(ID string, handler func()) (freq bool) {
	cnt := fa.access[ID]
	if cnt == 1 {
		fa.start = time.Now().Unix()
	}
	if cnt >= fa.maxCount {
		end := time.Now().Unix()
		if end-fa.start > int64(fa.maxDuration) {
			fmt.Println("freq: ")
			freq = true
			fa.access[ID] = 0
			if handler != nil {
				handler()
			}
			return
		}
	}
	return
}
