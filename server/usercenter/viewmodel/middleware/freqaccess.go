package usermiddleware

import (
	"time"
)

//FreqAccess 访问频度
type FreqAccess struct {
	maxCount    int             //最大统计次数
	maxDuration int             //最大时间限制
	start       int64           //开始时刻
	accesscnt   map[string]int  //当前统计次数
	freq        map[string]bool //是否频繁
	succ        map[string]bool //是否成功
}

//NewFreqAccess 初始化访问频度对象
func NewFreqAccess(maxDuration int, maxCount int) *FreqAccess {
	if maxDuration < 60 {
		maxDuration = 60
	}
	return &FreqAccess{
		maxDuration: maxDuration,
		maxCount:    maxCount,
		accesscnt:   make(map[string]int),
		freq:        make(map[string]bool),
		succ:        make(map[string]bool),
	}
}

//SetAccessSucc 访问是否成功
func (fa *FreqAccess) SetAccessSucc(ID string, succ bool) {
	fa.succ[ID] = succ
}

//AccessCtrl 访问控制
func (fa *FreqAccess) AccessCtrl(ID string, handler func() (reset bool)) (freq bool) {
	if fa.freq[ID] || fa.succ[ID] {
		return
	}
	fa.accesscnt[ID]++
	if fa.accesscnt[ID] == 1 {
		fa.start = time.Now().Unix()
	}
	if fa.accesscnt[ID] >= fa.maxCount { //请求次数超过限定次数
		end := time.Now().Unix()
		if end-fa.start < int64(fa.maxDuration) { //结束时间与开始时间比较
			fa.accesscnt[ID] = 0
			fa.freq[ID] = true
			freq = true
			if handler != nil {
				reset := handler()
				if reset {
					fa.freq[ID] = false
				}
			}
			return
		}
	}
	return
}
