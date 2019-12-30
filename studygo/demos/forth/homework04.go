package homework04

import (
	"fmt"
)

type Callback func(data interface{})

type SimpleGoRoutine struct {
	dataChan chan interface{}
}

/*创建简单goroutine*/
func CreateSimpleGoRoutine() *SimpleGoRoutine {
	return &SimpleGoRoutine{dataChan: make(chan interface{})}
}

/*写入数据*/
func (r *SimpleGoRoutine) Write(data interface{}) {
	if r == nil {
		return
	}
	r.dataChan <- data
}

/*读取数据*/
func (r *SimpleGoRoutine) Read(callback Callback) {
	go func() {
		for {
			select {
			case data := <-r.dataChan:
				callback(data)
			}
		}
	}()
}

/*事件压入事务池回调*/
type PoolExec func()

/*事务处理接口*/
type Job interface {
	Do()
}

/*事务处理对象*/
type Worker struct {
	JobQueue chan Job
}

/*创建事务处理实例*/
func CreateWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}

/*单个事务写入job*/
func (w Worker) WriteJob(job Job) {
	w.JobQueue <- job
}

/*单个事务执行job*/
func (w Worker) ExecJob(exec PoolExec) {
	go func() {
		for {
			exec()
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

/*执行单个事务处理实例处理事务*/
func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

/*事务处理池对象*/
type WorkerPool struct {
	workerLen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

/*创建事务处理池实例*/
func CreateWorkerPool(workerLen int) WorkerPool {
	wp := WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerLen),
	}
	for i := 0; i < workerLen; i++ {
		w := CreateWorker()
		w.ExecJob(func() {
			wp.WorkerQueue <- w.JobQueue
		})
	}
	return wp
}

/*事务池写入事务*/
func (wp WorkerPool)WriteJob(job Job)  {
	wp.JobQueue <- job
}

/*事务池分发事务处理对象执行事务*/
func (wp WorkerPool) ExecWorker() {
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				w := <- wp.WorkerQueue
				w <- job
			}
		}
	}()
}

type Score struct {
	Num int
}

var Cnt int

func (s *Score) Do() {
	fmt.Println("num is: ", s.Num)
	Cnt++
	//time.Sleep(time.Second * 1)
}
