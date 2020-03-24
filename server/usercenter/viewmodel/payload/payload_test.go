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
	send     int
	sendChan chan int
}

func (cj chanJob) Do() {
	fmt.Println("send: ", cj.send)
	cj.sendChan <- cj.send
}

func TestChannelWorkerPool(t *testing.T) {
	cj := chanJob{
		send:     0,
		sendChan: make(chan int),
	}

	workerLen := 100
	p := CreateWorkerPool(workerLen)
	p.Run()
	for i := 0; i < 100; i++ {
		cj.send = i
		p.Write(cj)
	}
	for {
		select {
		case s := <-cj.sendChan:
			fmt.Println("send: ", s)
		}
	}
}

func BenchmarkChannelWorkerPool(t *testing.B) {
	cj := chanJob{
		send: 0,
	}

	workerLen := 100
	p := CreateWorkerPool(workerLen)
	p.Run()
	for i := 0; i < t.N; i++ {
		cj.send = i
		p.Write(cj)
	}
	for {
		select {
		case s := <-cj.sendChan:
			fmt.Println("send: ", s)
		}
	}
}
