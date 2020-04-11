package userpayload

//https://blog.51cto.com/11140372/2342953?source=dra
//Pool 事务池
// var Pool WorkerPool

// func init() {
// 	workLen := 100
// 	Pool = CreateWorkerPool(workLen)
// 	Pool.Run()
// }

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
				go func(j Job) {
					job.Do()
				}(job)
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

//WorkerPool 事务池
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

//WriteHandler 写入事务池handler
func (p WorkerPool) WriteHandler(handler func(j chan Job)) {
	go func() {
		if handler != nil {
			handler(p.JobQueue)
		}
	}()
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
