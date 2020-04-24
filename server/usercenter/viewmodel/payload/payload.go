package userpayload

import (
	"os"
)

//https://blog.51cto.com/11140372/2342953?source=dra
//Pool 事务池
// var Pool WorkerPool

// func init() {
// 	workLen := 100
// 	Pool = NewWorkerPool(workLen)
// 	Pool.Run()
// }

var (
	//MaxWorker 最多事务数
	MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue 最大队列数
	MaxQueue = os.Getenv("MAX_QUEUE")
)

//Job 执行事务接口
type Job interface {
	Do()
}

//worker 事务对象
type worker struct {
	jobQueue chan Job
	quit     chan bool
}

//newWorker 创建事务
func newWorker() *worker {
	return &worker{
		jobQueue: make(chan Job),
		quit:     make(chan bool),
	}
}

//Run 执行事务
func (w *worker) run(wp chan chan Job) {
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
func (w *worker) stop() {
	go func() {
		w.quit <- true
	}()
}

//WorkerPool 事务池
type WorkerPool struct {
	maxWorker   int
	workerQueue chan chan Job
	jobQueue    chan Job
	quit        chan bool
}

//NewWorkerPool 创建事务池
func NewWorkerPool(maxWorker int) *WorkerPool {
	return &WorkerPool{
		maxWorker:   maxWorker,
		workerQueue: make(chan chan Job, maxWorker),
		jobQueue:    make(chan Job),
		quit:        make(chan bool),
	}
}

//Run 执行事务
func (p *WorkerPool) Run() {
	//初始化worker
	for i := 0; i < p.maxWorker; i++ {
		wk := newWorker()
		wk.run(p.workerQueue)
	}
	go p.dispatch()
}

//Stop 停止事务池
func (p *WorkerPool) Stop() {
	go func() {
		p.quit <- true
	}()
}

//WriteHandler 写入事务池handler
func (p *WorkerPool) WriteHandler(handler func(j chan Job)) {
	go func(jq chan Job) {
		if handler != nil {
			handler(p.jobQueue)
		}
	}(p.jobQueue)
}

//dispatch 事务分发
func (p *WorkerPool) dispatch() {
	for {
		select {
		case j := <-p.jobQueue:
			go func(job Job) {
				wk := <-p.workerQueue
				wk <- job
			}(j)
		case <-p.quit:
			return
		}
	}
}

/*
package Concurrence

import "fmt"

// --------------------------- Job ---------------------
type Job interface {
    Do()
}

// --------------------------- Worker ---------------------
type Worker struct {
    JobQueue chan Job
}

func NewWorker() Worker {
    return Worker{JobQueue: make(chan Job)}
}
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

// --------------------------- WorkerPool ---------------------
type WorkerPool struct {
    workerlen   int
    JobQueue    chan Job
    WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
    return &WorkerPool{
        workerlen:   workerlen,
        JobQueue:    make(chan Job),
        WorkerQueue: make(chan chan Job, workerlen),
    }
}
func (wp *WorkerPool) Run() {
    fmt.Println("初始化worker")
    //初始化worker
    for i := 0; i < wp.workerlen; i++ {
        worker := NewWorker()
        worker.Run(wp.WorkerQueue)
    }
    // 循环获取可用的worker,往worker中写job
    go func() {
        for {
            select {
            case job := <-wp.JobQueue:
                worker := <-wp.WorkerQueue
                worker <- job
            }
        }
    }()
}

// --------------- 使用 --------------------
/*
type Score struct {
    Num int
}

func (s *Score) Do() {
    fmt.Println("num:", s.Num)
    time.Sleep(1 * 1 * time.Second)
}

func main() {
    num := 100 * 100 * 20
    // debug.SetMaxThreads(num + 1000) //设置最大线程数
    // 注册工作池，传入任务
    // 参数1 worker并发个数
    p := NewWorkerPool(num)
    p.Run()
    datanum := 100 * 100 * 100 * 100
    go func() {
        for i := 1; i <= datanum; i++ {
            sc := &Score{Num: i}
            p.JobQueue <- sc
        }
    }()

    for {
        fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
        time.Sleep(2 * time.Second)
    }

}
package main
import "your/path/to/.../Concurrence"

//定义一个实现Job接口的数据
type Score struct {
    Num int
}
//定义对数据的处理
func (s *Score) Do() {
    fmt.Println("num:", s.Num)
    time.Sleep(1 * 1 * time.Second)
}

func main() {
    num := 100 * 100 * 20
    // debug.SetMaxThreads(num + 1000) //设置最大线程数
    // 注册工作池，传入任务
    // 参数1 worker并发个数
    p := NewWorkerPool(num)
    p.Run()

    //写入一亿条数据
    datanum := 100 * 100 * 100 * 100
    go func() {
        for i := 1; i <= datanum; i++ {
            sc := &Score{Num: i}
            p.JobQueue <- sc //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
        }
    }()

//循环打印输出当前进程的Goroutine 个数
    for {
        fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
        time.Sleep(2 * time.Second)
    }

}
*/

/*
var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool  chan chan Job
	JobChannel  chan Job
	quit    	chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					log.Errorf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
    // starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.pool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
*/
