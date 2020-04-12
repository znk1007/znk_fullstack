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
	freq        bool           //是否频繁
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
	if fa.freq {
		return
	}
	fa.access[ID]++
	if fa.access[ID] == 1 {
		fa.start = time.Now().Unix()
	}
	fmt.Println("cnt: ", fa.access[ID])
	if fa.access[ID] >= fa.maxCount { //请求次数超过限定次数
		end := time.Now().Unix()
		fmt.Println("diff: ", end-fa.start)
		if end-fa.start > int64(fa.maxDuration) { //结束时间与开始时间比较
			fa.access[ID] = 0
			if handler != nil {
				handler()
				fa.freq = true
			}
			return
		}
	}
	return
}
