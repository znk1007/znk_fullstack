package main

import (
	"fmt"

	"github.com/znk_fullstack/server/usercenter/viewmodel/payload"
)

func main() {
	workerLen := 100 * 100 * 100
	p := payload.CreateWorkerPool(workerLen)
	p.Run()
	dataNum := 100 * 100 * 100 * 100
	for i := 0; i < dataNum; i++ {
		j := simpleJob{
			num: i,
		}
		p.Write(j)
	}
}

type simpleJob struct {
	num int
}

func (s simpleJob) Do() {
	fmt.Println("job do, num: ", s.num)
}
