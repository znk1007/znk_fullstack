package homework04

type Callback func(data interface{})

type SimpleGoRoutine struct {
	dataChan chan interface{}
}
/*创建简单goroutine*/
func CreateSimpleGoRoutine() *SimpleGoRoutine {
	return &SimpleGoRoutine{dataChan:make(chan interface{})}
}
/*写入数据*/
func (r *SimpleGoRoutine) Write(data interface{}) {
	if r == nil {
		return
	}
	r.dataChan <- data
}
/*读取数据*/
func (r* SimpleGoRoutine)Read(callback Callback)  {
	go func() {
		for {
			select {
				case data := <- r.dataChan:
					callback(data)
			}
		}
	}()
}



type GoRoutineWorker struct {

}