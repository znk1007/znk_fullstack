package userpayload

//Pool 事务池
var Pool WorkerPool

func init() {
	workLen := 100
}

//Job 执行事务接口
type Job interface {
	Do()
}

//Worker 事务对象
type Worker struct {
	jobQueue chan Job
	quit     chan bool
}

//CreateWorker 创建事务
func CreateWorker() Worker {
	return Worker{
		jobQueue: make(chan Job),
		quit:     make(chan bool),
	}
}

//Write 写入事务
func (w Worker) Write(job Job) {
	w.jobQueue <- job
}

//Run 执行事务
func (w Worker) Run(wp chan chan Job) {
	go func() {
		for {
			wp <- w.jobQueue //regist current job channel to worker pool
			select {
			case job := <-w.jobQueue:
				job.Do()
			case <-w.quit:
				return
			}
		}
	}()
}

//Stop 停止事务
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type WorkerPool struct {
	maxWorker   int
	WorkerQueue chan chan Job
	JobQueue    chan Job
	quit        chan bool
}

//CreateWorkerPool 创建事务池
func CreateWorkerPool(maxWorker int) WorkerPool {
	return WorkerPool{
		maxWorker:   maxWorker,
		WorkerQueue: make(chan chan Job),
		JobQueue:    make(chan Job),
		quit:        make(chan bool),
	}
}

//Run 执行事务
func (p WorkerPool) Run() {
	for i := 0; i < p.maxWorker; i++ {
		wk := CreateWorker()
		wk.Run(p.WorkerQueue)
	}
	go p.dispatch()
}

//Stop 停止事务池
func (p WorkerPool) Stop() {
	go func() {
		p.quit <- true
	}()
}

//Write 写入事务池
func (p WorkerPool) Write(job Job) {
	p.JobQueue <- job
}

//dispatch 事务分发
func (p WorkerPool) dispatch() {
	for {
		select {
		case j := <-p.JobQueue:
			go func(job Job) {
				wk := <-p.WorkerQueue
				wk <- job
			}(j)
		case <-p.quit:
			return
		}
	}
}
