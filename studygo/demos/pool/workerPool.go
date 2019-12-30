package pool

import "fmt"

type Job interface {
	Do()
}

/*事务对象*/
type Worker struct {
	workerPool chan chan Job
	jobQueue   chan Job
	quit       chan bool
}
/*创建事务实例*/
func CreateWorker(wp chan chan Job) Worker {
	return Worker{
		workerPool: wp,
		jobQueue:   make(chan Job),
		quit:       make(chan bool),
	}
}

/*写入事务*/
func (w Worker) Write(job Job) {
	w.jobQueue <- job
}

/*处理事务*/
func (w Worker) Run() {
	go func() {
		for {
			if w.workerPool != nil {
				// 如果设置事务池，这把当前事务加入池中
				w.workerPool <- w.jobQueue
			}
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
		WorkerQueue: make(chan chan Job),
		JobQueue:    make(chan Job),
	}
	return wp
}
/*创建事务群并分发任务*/
func (wp WorkerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		fmt.Println("idx: ", i)
		work := CreateWorker(wp.WorkerQueue)
		work.Run()
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
			go func(job2 Job) {
				worker := <- wp.WorkerQueue
				worker <- job2
			}(job)
		}
	}
}

type Task struct {
	Num int
}

func (t Task) Do() {
	fmt.Println("task do, num: ", t.Num)
}
