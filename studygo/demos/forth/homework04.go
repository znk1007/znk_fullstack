package forth
//https://blog.csdn.net/Jeanphorn/article/details/79018205
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
func (w Worker) ExecJob() {
	go func() {
		for {
			select {
			case job := <-w.JobQueue:
				fmt.Println("w.JobQueue: ", job)
				job.Do()
			}
		}
	}()
}

/*事务池执行job*/
func (w Worker) ExecJobWithQueue(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue//注册当前job队列到事务池中
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
	WorkerChan  chan Worker
}

/*创建事务处理池实例*/
func CreateWorkerPool(workerLen int) WorkerPool {
	wp := WorkerPool{
		workerLen:   workerLen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerLen),
		WorkerChan:  make(chan Worker, workerLen),
	}
	for i := 0; i < workerLen; i++ {
		w := CreateWorker()
		w.ExecJobWithQueue(wp.WorkerQueue)
	}
	return wp
}

/*事务池写入事务*/
func (wp WorkerPool) WriteJob(job Job) {
	wp.JobQueue <- job

}

/*事务池分发事务处理对象执行事务*/
func (wp WorkerPool) ExecWorker() {
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				w := <-wp.WorkerQueue
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

//func main() {
//that := homework04.CreateSimpleGoRoutine()
//
//that.Read(func(data interface{}) {
//	fmt.Println("read simple data: ", data)
//})
//for i := 1; i <= 100; i++ {
//	that.Write(i)
//}
//	w := homework04.CreateWorker()
//	w.ExecJob()
//	dataNum := 100 * 100 * 100 * 100
//	for i := 1; i <= dataNum; i++ {
//		sc := &homework04.Score{Num:i}
//		//wp.ExecWorker(sc)
//		w.WriteJob(sc)
//	}
//	start := time.Now()
//	fmt.Println("start time: ", start)
//	num := 100 * 100 * 100
//	wp := homework04.CreateWorkerPool(num)
//	wp.ExecWorker()
//	dataNum := 100 * 100 * 100 //* 100
//	for i := 1; i <= dataNum; i++ {
//		sc := &homework04.Score{Num: i}
//		wp.WriteJob(sc)
//	}
//
//	fmt.Println("end time: ", time.Now().Second()-start.Second())
//}