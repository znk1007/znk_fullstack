package homework04

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
func (w Worker) WriteJob(job Job)  {
	w.JobQueue <- job
}

/*单个事务执行job*/
func (w Worker) ExecJob()  {
	go func() {
		for {
			select {
			case job := <- w.JobQueue:
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
	wokerLen    int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

/*创建事务处理池实例*/
func CreateWorkerPool(workerLen int) WorkerPool {
	return WorkerPool{
		wokerLen:    workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job),
	}
}

func (wp WorkerPool) Run() {

}
