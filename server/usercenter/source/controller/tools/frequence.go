package usertools

import (
	"fmt"
	"sync"
)

//Freq 频度
type Freq struct {
	acc     string
	method  string
	freqmap sync.Map
	expired int64
	ts      int64
}

//NewFreq 初始化
func NewFreq(expired int64) *Freq {
	return &Freq{
		expired: expired,
	}
}

//Expired 是否超时
func (f *Freq) Expired(acc, method string, ts int64) (expired bool) {
	key := acc + "_" + method
	org, _ := f.freqmap.Load(key)
	fmt.Println("org: ", org)
	if org != nil {
		orgTs, _ := org.(int64)
		if ts-orgTs > f.expired {
			expired = true
		}
	}
	f.freqmap.Store(key, ts)
	return
}
