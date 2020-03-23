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
	workerLen := 100 * 100 * 100
	p := CreateWorkerPool(workerLen)
	p.Run()

	for i := 0; i < t.N; i++ {
		j := simpleJob{
			num: i,
		}
		p.Write(j)
	}
}
