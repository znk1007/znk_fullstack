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
	workerLen := 100 * 100 * 10
	p := CreateWorkerPool(workerLen)
	p.Run()
	p.WriteHandler(func(jq chan Job) {
		for i := 0; i < t.N; i++ {
			j := simpleJob{
				num: i,
			}
			jq <- j
		}
	})

}

func TestSimpleWorkerPoolHandler(t *testing.T) {
	workerLen := 100 * 10
	p := CreateWorkerPool(workerLen)
	p.Run()
	num := 100 * 500
	p.WriteHandler(func(jq chan Job) {

		for i := 0; i < num; i++ {
			j := simpleJob{
				num: i,
			}
			jq <- j
		}
	})
}

func BenchmarkSimpleWorkerPoolHandler(t *testing.B) {
	workerLen := 100 * 10
	p := CreateWorkerPool(workerLen)
	p.Run()
	// num := 100 * 500
	p.WriteHandler(func(jq chan Job) {

		for i := 0; i < t.N; i++ {
			j := simpleJob{
				num: i,
			}
			jq <- j
		}
	})
}

type chanJob struct {
	send     int
	sendChan chan int
}

func (cj chanJob) Do() {
	go func() {
		fmt.Println("send: ", cj.send)
		cj.sendChan <- cj.send
	}()
}

func TestChannelWorkerPool(t *testing.T) {
	cj := chanJob{
		send:     0,
		sendChan: make(chan int),
	}

	workerLen := 1000
	p := CreateWorkerPool(workerLen)
	p.Run()
	p.WriteHandler(func(jq chan Job) {
		for i := 0; i < 100; i++ {
			cj.send = i
			jq <- cj
		}
	})
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
	p.WriteHandler(func(jq chan Job) {
		for i := 0; i < 2; i++ {
			cj.send = i
			jq <- cj
		}
	})
	for {
		select {
		case s := <-cj.sendChan:
			fmt.Println("send: ", s)
		}
	}
}
