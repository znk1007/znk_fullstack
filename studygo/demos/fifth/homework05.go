package fifth

import "fmt"
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
	testMap["name"] = Test{Name:"测试"}
	fmt.Println("name: ", testMap["name"])
	
	var testMap1 map[string]*Test  = map[string]*Test{}
	testMap1["name"] = &Test{Name:"测试一"}
	testMap1["name"].Name = "测试二"
	fmt.Println("name1: ", testMap1["name"])
}

const (
	a = iota//iota 换⾏值+1
	b
	c = "c"
	d = iota
)

func PrintConst() {
	fmt.Println(a, b, c, d)
}

func Foo(x interface{})  {
	fmt.Println(x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func test() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		funcs = append(funcs, func() {
			fmt.Println(&i, i)
		})
	}
	return funcs
}
// 闭包延迟求值 for循环复⽤局部变量i，每⼀次放⼊匿名函数的应⽤都是同⼀个变量。
func AppendFunc() {
	funcs := test()
	for _, f := range funcs {
		f()
	}
}