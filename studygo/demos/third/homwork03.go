package homework03

import "fmt"

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
	if index < 0 {
		that.insertHead(data)
	} else if index > that.size {
		that.insertTail(data)
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
