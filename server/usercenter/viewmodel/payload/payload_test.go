package userpayload

import (
	"fmt"
	"testing"
)

type simpleJob struct {
	num int
}

func (j simpleJob) Do() {
	fmt.Println("job do, num: ", j.num)
}

func BenchmarkSimpleWorkerPool(t *testing.B) {
	workerLen := 100 * 100 * 100
	p := CreateWorkerPool(workerLen)
	p.Run()

	for i := 0; i < t.N; i++ {
		j := simpleJob{
			num: i,
		}
		p.Write(j)
	}
}

type chanJob struct {
	sendChan chan int
}

func (cj chanJob) Do() {
	send := <-cj.sendChan
	fmt.Println("send: ", send)
}

func TestChannelWorkerPool(t *testing.T) {
	cj := chanJob{
		sendChan: make(chan int),
	}

	workerLen := 100
	p := CreateWorkerPool(workerLen)
	p.Run()
	for i := 0; i < 10000; i++ {
		cj.sendChan <- i
		p.Write(cj)
	}
}

func BenchmarkChannelWorkerPool(t *testing.B) {
	cj := chanJob{
		sendChan: make(chan int),
	}

	workerLen := 100
	p := CreateWorkerPool(workerLen)
	p.Run()
	for i := 0; i < t.N; i++ {
		cj.sendChan <- i
		p.Write(cj)
	}
}
