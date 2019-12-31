package pool

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


type Job interface {
	Do()
}

/*事务对象*/
type Worker struct {
	jobQueue   chan Job
	quit       chan bool
}
/*创建事务实例*/
func CreateWorker() Worker {
	return Worker{
		jobQueue:   make(chan Job),
		quit:       make(chan bool),
	}
}

/*写入事务*/
func (w Worker) Write(job Job) {
	w.jobQueue <- job
}

/*处理事务*/
func (w Worker) Run(wp chan chan Job) {
	go func() {
		for {
			// 这把当前事务加入池中
			wp <- w.jobQueue
			select {
			case job := <-w.jobQueue:
				job.Do()
			case <-w.quit:
				return
			}
		}
	}()
}
/*停止事务处理*/
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
/*事务池对象*/
type WorkerPool struct {
	maxWorker   int
	WorkerQueue chan chan Job
	JobQueue    chan Job
}
/*创建事务池*/
func CreateWorkerPool(maxWorker int) WorkerPool {
	wp := WorkerPool{
		maxWorker:   maxWorker,
		WorkerQueue: make(chan chan Job, maxWorker),
		JobQueue:    make(chan Job),
	}
	return wp
}
/*创建事务群并分发任务*/
func (wp WorkerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		work := CreateWorker()
		work.Run(wp.WorkerQueue)
	}
	go wp.dispatch()
	
}

/*写入事务*/
func (wp WorkerPool) Write(job Job) {
	wp.JobQueue <- job
}

/*事务分发给事务处理对象*/
func (wp WorkerPool) dispatch() {
	for {
		select {
		case job := <-wp.JobQueue:
			worker := <- wp.WorkerQueue
			worker <- job
		}
	}
}

type Task struct {
	Num int
}

func (t Task) Do() {
	fmt.Println("task do, num: ", t.Num)
}

//func main() {
//	dataNum := 100 * 100 * 100 * 100
//	test := true
//	if test {
//		start := time.Now()
//		fmt.Println("start time: ", start)
//		workLen := 100 * 100 * 100
//		wp := pool.CreateWorkerPool(workLen)
//		wp.Run()
//		for i := 0; i < dataNum; i++ {
//			t := pool.Task{Num:i}
//			wp.Write(t)
//		}
//
//		fmt.Println("end time: ", time.Now().Second()-start.Second())
//	} else {
//		w := pool.CreateWorker()
//		w.Run(nil)
//		for i := 0; i < dataNum; i++ {
//			t := pool.Task{Num:i}
//			w.Write(t)
//		}
//	}
//
//
//}
