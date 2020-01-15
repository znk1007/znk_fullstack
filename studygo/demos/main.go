package main

import (
	"github.com/znk_fullstack/studygo/demos/eighth"
	"time"
)

func main() {
	//list := chain.CreateLinkedList(true)
	//for i := 0; i < 12; i++ {
	//	list.InsertLinkedListNode(-1, i)
	//}
	////head := list.Head()
	////head.Print(true, true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++0++++++++++")
	//
	//list.InsertLinkedListNode(5, 20)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//node := list.Get(5)
	//fmt.Println("the node: ", node)
	//
	//fmt.Println("+++++++++++++1+++++++++++++")
	//
	//list.InsertLinkedListNode(8, 20)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++++2+++++++++++++")
	//idxs := list.Search(20)
	//fmt.Println("idxs: ", idxs)
	//
	//fmt.Println("+++++++++++++3+++++++++++++")
	//
	//succ := list.DeleteByIndex(5)
	//fmt.Println("delete succ: ", succ)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++++4+++++++++++++")
	//list.Destory()
	//list.Print(true)
	//fmt.Println("-------insert------ 1")
	//that := homework03.CreateLinkedList()
	//for i := 0; i < 10; i++ {
	//	that.Insert(i, -1)
	//}
	//that.Print()
	//
	//fmt.Println("--------1------- length: ", that.Length())
	//
	//fmt.Println("-------insert------ 2")
	//
	//for i := 10; i < 20; i++ {
	//	that.Insert(i, that.Length()+1)
	//}
	//that.Print()
	//
	//fmt.Println("--------2------- length: ", that.Length())
	//
	//
	//fmt.Println("--------2------- is circular: ", that.IsCircular())
	//
	//fmt.Println("--------3------- length: ", that.Length())
	//fmt.Println("tail: ", that.Tail())
	//node := that.Get(10)
	//fmt.Println("get node: ", node)
	//
	//fmt.Println("--------4-------")
	//for i := 20; i < 30; i++ {
	//	that.Insert(i, 2)
	//}
	//
	//that.Print()

	//fmt.Println("--------5-------")
	//
	//for i := 1; i < 11; i++ {
	//	that.Insert(i, that.Length()+1)
	//}
	//
	//that.Print()
	//fmt.Println("--------5------- is circular: ", that.IsCircular())
	//fmt.Println("--------5------- length: ", that.Length())
	//
	//fmt.Println("--------6------- delete: ")
	//
	//that.DeleteByIndex(2)
	//
	//that.Print()
	//fmt.Println("--------6------- is circular: ", that.IsCircular())
	//fmt.Println("--------6------- length: ", that.Length())
	//
	//fmt.Println("--------7-------")
	//
	//node := that.Search(10, true)
	//fmt.Println("search node: ", node)
	//
	//fmt.Println("--------8------- delete")
	//
	//that.DeleteByData(10, true)
	//that.Print()
	//fmt.Println("--------8------- head: ", that.Head())
	//fmt.Println("--------8------- tail: ", that.Tail())
	//fmt.Println("--------8------- is circular: ", that.IsCircular())
	//fmt.Println("--------8------- length: ", that.Length())

	//that := homework03.CreateJosephus()
	//that.Insert(1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	//	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	//	21, 22, 23, 24, 25, 26, 27, 28, 28, 30,
	//	31, 32, 33, 34, 35, 36, 37, 38, 38, 40,
	//	41,
	//)
	////that.Print()
	//fmt.Println("length: ", that.Length())
	//that.Escape()
	//
	//ch := make(chan int)
	//
	//go func() {
	//	for val := range ch {
	//		fmt.Println("value 1: ", val)
	//	}
	//}()
	//ch <- 100
	//
	//ch = make(chan int, 1CreateWorker)
	//ch <- 123
	//close(ch)
	//value := <-ch
	//fmt.Println("value 2: ", value)
	//
	//var m sync.Mutex
	//c := 0
	//wait := sync.WaitGroup{}
	//for i := 0; i < 100; i++ {
	//	wait.Add(1)
	//	go func() {
	//		m.Lock()
	//		defer m.Unlock()
	//		wait.Done()
	//		//fmt.Println("current c: ", c)
	//		c++
	//	}()
	//}
	//
	//wait.Wait()
	//fmt.Println("c:", c)

	//that := homework04.CreateSimpleGoRoutine()
	//
	//that.Read(func(data interface{}) {
	//	fmt.Println("read simple data: ", data)
	//})
	//for i := 1; i <= 100; i++ {
	//	that.Write(i)
	//}

	//w := homework04.CreateWorker()
	//w.ExecJob()
	//dataNum := 100 * 100 * 100 * 100
	//for i := 1; i <= dataNum; i++ {
	//	sc := &homework04.Score{Num:i}
	//	//wp.ExecWorker(sc)
	//	w.WriteJob(sc)
	//}
	//test := false
	//if test {
	//	start := time.Now()
	//	fmt.Println("start time: ", start)
	//	num := 100 * 100 * 100
	//	wp := forth.CreateWorkerPool(num)
	//	wp.ExecWorker()
	//	dataNum := 100 * 100 * 100 //* 100
	//	for i := 1; i <= dataNum; i++ {
	//		sc := &forth.Score{Num: i}
	//		wp.WriteJob(sc)
	//	}
	//
	//	fmt.Println("end time: ", time.Now().Second()-start.Second())
	//} else {
	//
	//	dataNum := 100 * 100 * 100 //* 100
	//	//w := pool.CreateWorker()
	//	//w.Run(nil)
	//	//for i := 0; i < dataNum; i++ {
	//	//	t := pool.Task{Num:i}
	//	//	w.Write(t)
	//	//}
	//	start := time.Now()
	//	fmt.Println("start time: ", start)
	//	workLen := 100 * 100 * 100
	//	wp := pool.CreateWorkerPool(workLen)
	//	wp.Run()
	//	for i := 0; i < dataNum; i++ {
	//		t := pool.Task{Num:i}
	//		wp.Write(t)
	//	}
	//
	//	fmt.Println("end time: ", time.Now().Second()-start.Second())
	//}

	//fifth.Array()
	//fifth.AppendSlice()
	//fifth.MapTest()
	//fifth.PrintConst()
	//
	//var x *int  = nil
	//fifth.Foo(x)
	//
	//fifth.AppendFunc()
	//
	//fifth.ReverseStr()
	//
	//var stu = fifth.Student{}
	//var peo fifth.People = &fifth.Student{}
	//stuStr := stu.Speak("法师")
	//fmt.Println("stu str: ", stuStr)
	//peoStr := peo.Speak("法师")
	//fmt.Println("peo str: ", peoStr)
	//translation := seventh.Translation{
	//	File:"seventh/dict.txt",
	//}
	//translation.Translate("嫣然")
	//fmt.Println("src: ", translation.Src)
	//fmt.Println("val: ", translation.Target)
	// trans := seventh.CreateTranslation("seventh/dict.txt")
	// trans.GetResult(func(src string, result string) {
	// 	fmt.Println("src1: ", src)
	// 	fmt.Println("result1: ", result)
	// })
	// trans.Translate("嫣然")
	// trans.Translate2("嫣然", func(src string, target string) {
	// 	fmt.Println("src2: ", src)
	// 	fmt.Println("result2: ", target)
	// })

	//src -> /Users/huangsam/Downloads/Boom_40054.zip
	//dst -> /Users/huangsam/Downloads/Boom_40051.zip

	// go run main.go /Users/huangsam/Downloads/Boom_40054.zip /Users/huangsam/Downloads/Boom_40051.zip

	// args := os.Args
	// if args != nil && len(args) >= 3 {
	// 	srcPath := args[1] //获取输⼊的第⼀个参数
	// 	dstPath := args[2] //获取输⼊的第⼆个参数
	// 	fmt.Printf("srcPath = %s, dstPath = %s\n", srcPath, dstPath)
	// 	if srcPath != dstPath {
	// 		c := seventh.CreateCopy(srcPath, dstPath)
	// 		c.Copy()
	// 	}

	// }
	srcPath := "/Users/huangman/Desktop/golang/src/github.com/znk_fullstack/studygo/demos/eighth/Go语言工程师信息表.xlsx"
	dstPath := "/Users/huangman/Desktop/golang/src/github.com/znk_fullstack/studygo/demos/eighth/Go语言工程师信息表-New.xlsx"
	codeInfo := eighth.CreateCoderInfo(srcPath, dstPath)
	codeInfo.ReadAndSave("1-3年")

	time.Sleep(time.Millisecond * 500)
}
