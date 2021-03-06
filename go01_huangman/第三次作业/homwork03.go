package homework03

import (
	"fmt"
	"reflect"
	"sync"
)

/*链表对象*/
type LinkedList struct {
	 size int
	 head *LinkedListNode
	 tail *LinkedListNode
}
/*链表节点对象*/
type LinkedListNode struct {
	data interface{}
	prev *LinkedListNode
	next *LinkedListNode
}
/*创建链表对象*/
func CreateLinkedList() *LinkedList  {
	return &LinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}
/*打印链表*/
func (that *LinkedList)Print()  {
	if that == nil {
		return
	}
	node := that.head
	for node != that.tail {
		fmt.Println("node data: ", node.data)
		node = node.next
	}
	fmt.Println("node data: ", node.data)
}
/*链表长度*/
func (that *LinkedList) Length() int {
	if that == nil {
		return 0
	}
	return that.size
}

func (that *LinkedList)	Insert(data interface{}, index int)  {
	if index < 1 {
		that.insertHead(data)
	} else if index > that.size {
		that.insertTail(data)
	} else {
		newNode := &LinkedListNode{
			data: data,
			prev: nil,
			next: nil,
		}
		if that.head == nil {
			that.head = newNode
			that.tail = newNode
			that.size++
		} else {
			node := that.Get(index)
			if node == nil {
				return
			}
			newNode.next = node.next
			node.next = newNode
			that.size++
		}
	}
}

/*插入头部*/
func (that *LinkedList) insertHead(data interface{})  {
	if that == nil || data == nil {
		return
	}
	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if that.head == nil {
		 that.tail = newNode
	}
	newNode.next = that.head
	that.head = newNode
	that.tail.next = that.head
	that.size++
}
/*插入尾部*/
func (that *LinkedList) insertTail(data interface{})  {
	if that == nil || data == nil {
		return
	}
	newNode := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if that.head == nil {
		that.head = newNode
		that.tail = newNode
		that.size++
	} else {
		that.tail.next = newNode
		that.tail = newNode
		that.tail.next = that.head
		that.size++
	}
}
/*获取下标节点*/
func (that *LinkedList) Get(index int) *LinkedListNode {
	if that == nil || index < 1 || index > that.size || that.size == 0 {
		return nil
	}
	if index == 1 {
		return that.head
	}
	if index == that.size {
		return that.tail
	}
	node := that.head
	idx := 1
	for node != that.tail {
		idx++
		node = node.next
		if idx == index {
			return node
		}
	}
	return nil
}
/*获取数据所在节点下标*/
func (that *LinkedList) Search(data interface{}, once bool) []int  {
	var idxes []int
	if that == nil || data == nil  || that.size == 0 {
		return idxes
	}
	node := that.head
	idx := 1
	for node != that.tail  {
		if reflect.TypeOf(data) == reflect.TypeOf(node.data) && data == node.data {
			idxes = append(idxes, idx)
			if once {
				return idxes
			}
		}
		idx++
		node = node.next
	}
	if len(idxes) != 0 && once {
		return idxes
	}
	idx++
	if reflect.TypeOf(data) == reflect.TypeOf(node.data) && data == node.data {
		idxes = append(idxes, idx)
	}
	return idxes
}
/*根据下标删除*/
func (that *LinkedList) DeleteByIndex(index int) {
	if that == nil || index < 1 || index > that.size || that.size == 0 {
		return
	}
	node := that.head
	if that.size == 1 {
		 that.head = nil
		 that.tail = nil
		 that.size = 0
	} else if index == 1 {
		that.head = node.next
		node.next = nil
		node.data = nil
		node = nil
		that.size--
	} else if index == that.size {
		prev := that.head
		for node != that.tail {
			prev = node
			node = node.next
		}
		that.tail = prev
		that.tail.next = that.head
		node.next = nil
		node.data = nil
		node = nil
		that.size--
	} else {
		prev := that.head
		for i := 2; i <= index; i++ {
			prev = node
			node = node.next
		}
		prev.next = node.next
		node.next = nil
		node.data = nil
		node = nil
		that.size--
	}
}
/*根据数据删除*/
func (that *LinkedList) DeleteByData(data interface{}, once bool) {
	if that == nil || data == nil || that.size == 0 {
		return
	}
	var del bool = false
	node := that.head
	prev := node
	if that.size == 1 && reflect.TypeOf(data) == reflect.TypeOf(that.head.data) && data == that.head.data {
		that.head = nil
		that.tail = nil
		that.size--
	} else {
		for node != that.tail {
			if reflect.TypeOf(data) == reflect.TypeOf(that.head.data) && data == that.head.data {
				that.tail.next = node.next
				that.head = node.next
				node.next = nil
				node.data = nil
				node = nil
				that.size--
				del = true
				if once {
					return
				}
			} else {
				if reflect.TypeOf(data) == reflect.TypeOf(node.data) && data == node.data {
					prev.next = node.next
					node.next = nil
					node.data = nil
					node = nil
					del = true
					that.size--
					if once {
						return
					}
				}
			}
			prev = node
			node = node.next
		}
	}

	if once && del {
		return
	}

	if that.tail != nil && reflect.TypeOf(data) == reflect.TypeOf(that.tail.data) && data == that.tail.data {
		prev.next = node.next
		that.tail = prev
		that.tail.next = that.head
		node.next = nil
		node.data = nil
		node = nil
		that.size--
	}
}

/*头结点*/
func (that *LinkedList) Head() *LinkedListNode {
	if that == nil {
		return nil
	}
	return that.head
}
/*尾结点*/
func (that *LinkedList) Tail() *LinkedListNode {
	if that == nil {
		return nil
	}
	return that.tail
}
/*是否循环*/
func (that *LinkedList) IsCircular() bool {
	if that == nil {
		return false
	}
	return that.tail.next == that.head
}

type input struct {
	lhs float64
	rhs float64
}

func (ip input) result(ch string) float64 {
	switch ch {
	case "+":
		return ip.lhs + ip.rhs
	case "-":
		return ip.lhs - ip.rhs
	case "*":
		return ip.lhs * ip.rhs
	case "/":
		if ip.rhs == 0 {
			return -1
		}
		return ip.lhs / ip.rhs
	}
	return 0
}

type calculate interface {
	result() float64
}

type add struct {
	input
}

func (a add)result() float64 {
	return a.lhs + a.rhs
}

type sub struct {
	input
}

func (s sub)result() float64 {
	return s.lhs - s.rhs
}

type mul struct {
	input
}

func (m mul) result() float64 {
	return m.lhs * m.rhs
}

type div struct {
	input
}
func (d div) result() float64  {
	if d.rhs == 0 {
		return -1
	}
	return d.lhs / d.rhs
}

func Calculate01(lhs float64, rhs float64, ch string) float64 {
	i := input{
		lhs: lhs,
		rhs: rhs,
	}
	return i.result(ch)
}

func Calculate02(lhs float64, rhs float64, ch string) float64 {
	var cal calculate
	switch ch {
	case "+":
		cal = add{input{
			lhs: lhs,
			rhs: rhs,
		}}
	case "-":
		cal = sub{input{
			lhs: lhs,
			rhs: rhs,
		}}
	case "*":
		cal = mul{input{
			lhs: lhs,
			rhs: rhs,
		}}
	case "/":
		cal = div{input{
			lhs: lhs,
			rhs: rhs,
		}}
	}
	return cal.result()
}


/*约瑟夫环*/
type Josephus struct {
	mutex *sync.RWMutex
	head *JosephusNode
	tail *JosephusNode
	size int
}
/*约瑟夫环*/
type JosephusNode struct {
	data interface{}
	next *JosephusNode
}
/*创建约瑟夫环*/
func CreateJosephus() *Josephus {
	return &Josephus{
		mutex:new(sync.RWMutex),
		head: nil,
		tail:nil,
		size: 0,
	}
}

func (that *Josephus) Insert(datas ...interface{}) bool {
	if that == nil || len(datas) == 0 {
		return false
	}
	that.mutex.Lock()
	defer that.mutex.Unlock()
	for _, data := range datas {
		newNode := &JosephusNode{
			data: data,
			next: nil,
		}
		if that.head == nil {
			that.head = newNode
			that.tail = newNode
		} else {
			that.tail.next = newNode
			that.tail = newNode
		}
		that.size++
	}
	that.tail.next = that.head
	return true
}
/*打印约瑟夫环*/
func (that *Josephus) Print() {
	if that == nil || that.size == 0 {
		return
	}
	node := that.head
	for node != that.tail {
		fmt.Println("node data: ", node.data)
		node = node.next
	}
	fmt.Println("node data: ", node.data)
}
/*约瑟夫环长度*/
func (that *Josephus) Length() int {
	if that == nil {
		return 0
	}
	return that.size
}
/*逃命*/
func (that *Josephus) Escape() {
	if that == nil || that.size == 0 {
		return
	}
	n := that.size
	m := 3
	node := that.head // 第一节点
	tempNode := that.head
	m %= n // 2
	for node != node.next {
		for i := 1; i < m - 1; i++ {
			node = node.next // 第二节点
		}
		fmt.Printf("%v->", node.next.data) // n / 3 == 0节点数据
		tempNode = node.next // 需要删除的 n / 3 == 0 节点
		node.next = tempNode.next
		tempNode.data = nil
		tempNode.next = nil
		tempNode = nil
		node = node.next
	}
	fmt.Printf("%v ", node.data)
	fmt.Println("head: ", that.head)
	fmt.Println("tail: ", that.tail)

}


