package fifth

import (
	"fmt"
	"sync"
)

//s = s[low : high : max] 切片的三个参数的切片截取的意义为
//low为截取的起始下标（含），
//high为窃取的结束下标（不含high），
//max为切片保留的原切片的最大下标（不含max）；
//即新切片从老切片的low下标元素开始，
//len = high - low, cap = max - low；
//high 和 max一旦超出在老切片中越界，
//就会发生runtime err，slice out of range。
//另外如果省略第三个参数的时候，第三个参数默认和第二个参数相同，即len = cap。

func Array() {
	msg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sli1 := msg[2:3:4]
	sli2 := msg[2:3]
	fmt.Println("sli1, len, cap", sli1, len(sli1), cap(sli1))
	fmt.Println("sli2, len, cap", sli2, len(sli2), cap(sli2))

}

func AppendSlice() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println("s1: ", s1)
}

type Test struct {
	Name string
}

/*
list[“name”]不是⼀个普通的指针值，map的value本身是不可寻址的，因为map中的值会在内存
中移动，并且旧的指针地址在map改变时会变得⽆效。 定义的是var list map[string]Test，注意
哦Test不是指针，⽽且map我们都知道是可以⾃动扩容的，那么原来的存储name的Test可能在地
址A，但是如果map扩容了地址A就不是原来的Test了，所以go就不允许写数据。改为var list
map[string]*Test。
*/
func MapTest() {
	var testMap map[string]Test = map[string]Test{}
	testMap["name"] = Test{Name: "测试"}
	fmt.Println("name: ", testMap["name"])

	var testMap1 map[string]*Test = map[string]*Test{}
	testMap1["name"] = &Test{Name: "测试一"}
	testMap1["name"].Name = "测试二"
	fmt.Println("name1: ", testMap1["name"])
}

const (
	a = iota //iota 换⾏值+1
	b
	c = "c"
	d = iota
)

func PrintConst() {
	fmt.Println(a, b, c, d)
}

func Foo(x interface{}) {
	fmt.Println(x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func test() ([]func(), []func()) {
	var funcs1 []func()
	for i := 0; i < 2; i++ {
		funcs1 = append(funcs1, func() {
			fmt.Println(&i, i)
		})
	}
	var funcs2 []func()
	for i := 0; i < 2; i++ {
		i := i
		funcs2 = append(funcs2, func() {
			fmt.Println(&i, i)
		})
	}
	return funcs1, funcs2
}

// 闭包延迟求值 for循环复⽤局部变量i，每⼀次放⼊匿名函数的应⽤都是同⼀个变量。
func AppendFunc() {
	funcs1, funcs2 := test()
	for _, f := range funcs1 {
		f()
	}
	for _, f := range funcs2 {
		f()
	}
}

func reverse(str string) string {
	rs := []rune(str)
	strLen := len(rs)

	var tt = make([]rune, 0)
	for i := 0; i < strLen; i++ {
		tt = append(tt, rs[strLen-i-1])
	}
	return string(tt[0:])
}

func reverse1(str string) string {
	rs := []rune(str)
	strLen := len(rs)
	for i := 0; i < strLen/2; i++ {
		rs[i], rs[strLen-1-i] = rs[strLen-1-i], rs[i]
	}
	return string(rs)
}

func ReverseStr() {
	str := "锄⽲⽇当午"
	fmt.Println("reverse: ", reverse(str))
	str = "锄⽲⽇当午，汗滴禾下土"
	fmt.Println("reverse1: ", reverse1(str))
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "法师" {
		talk = "法师，我爱你哟～"
	} else {
		talk = "hi"
	}
	return
}

func DeferPrint() {
	var name = "zhangsan"
	fmt.Println(name)
	defer fmt.Println(name)
	name = "lisa"
	defer fmt.Println(name)
}

type Singleton struct {

}

var single *Singleton
var mu sync.Mutex
// //懒汉加锁:虽然解决并发的问题，但每次加锁是要付出代价的
func GetSingleton() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if single == nil {
		 single = &Singleton{}
	}
	return single
}

func GetSingleton1() *Singleton {
	if single == nil {
		 mu.Lock()
		 defer mu.Unlock()
		if single == nil {
			single = &Singleton{}
		}
	}
	return single
}

var once sync.Once

func GetSingleton2() *Singleton {
	once.Do(func() {
		single = &Singleton{}
	})
	return single
}

/*

列举golang协程间通信常⻅的⼏种⽅法
1.使⽤channel机制，每个goroutine传⼀个channel进去然后往⾥写数据，在再主线程中读取这些
channel，直到全部读到数据了⼦goroutine也就全部运⾏完了，那么主goroutine也就可以结束
了。这种模式是⼦线程去通知主线程结束。
2.使⽤context中cancel函数，这种模式是主线程去通知⼦线程结束。
3.sync.WaitGroup模式，Add⽅法设置等待⼦goroutine的数量，使⽤Done⽅法设置等待⼦
goroutine的数量减1，当等待的数量等于0时，Wait函数返回。
*/