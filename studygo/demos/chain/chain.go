package chain

import (
	"fmt"
	"reflect"
)

/*链表对象*/
type LinkedList struct {
	size int
	single bool
	head *LinkedListNode
}

/*链表节点对象*/
type LinkedListNode struct {
	 data interface{}
	 prev *LinkedListNode
	 next *LinkedListNode
}

/*创建链表*/
func CreateLinkedList(single bool) *LinkedList  {
	linkedList := &LinkedList{
		size:   0,
		single: single,
		head:   nil,
	}
	return linkedList
}

/*插入节点*/
func (self *LinkedList)InsertLinkedListNode(index int, data interface{}) bool  {
	if self == nil || data == nil {
		return false
	}
	if index < 0 {
		self.insertHead(data)
	} else if index > self.size {
		self.insertTail(data)
	} else {
		node := self.Get(index - 1)
		if node == nil {
			return false
		}
		newNode := &LinkedListNode{
			data: data,
			prev: nil,
			next: nil,
		}
		newNode.next = node.next
		node.next = newNode
		if self.single == false {
			newNode.prev = node
		}
		self.size++
	}
	return false
}
/*链表长度*/
func (self *LinkedList) Length() int {
	return self.size
}

/*打印链表数据*/
func (self *LinkedList)Print(currentOnly bool)  {
	self.head.print(currentOnly)
}

/*打印链表数据*/
func (node *LinkedListNode)print(currentOnly bool)  {
	if node == nil {
		return
	}
	if currentOnly {
		fmt.Println("current node data: ", node.data)
	} else {
		fmt.Println("----start----")
		if node.prev != nil {
			fmt.Println("prev node data: ", node.prev.data)
		} else {
			fmt.Println("prev node is nil")
		}
		fmt.Println("current node data: ", node.data)
		if node.next != nil {
			fmt.Println("next node data: ", node.next.data)
		} else {
			fmt.Println("next node is nil")
		}
		
		fmt.Println("----end----")
	}
	node.next.print(currentOnly)
}

/*头部节点*/
func (self *LinkedList) Head() *LinkedListNode  {
	return self.head
}

/*插入头部*/
func (self *LinkedList)insertHead(data interface{})  {
	if self == nil || data == nil {
		return
	}
	node := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if self.single {
		node.next = self.head
		self.head = node
		self.size++
	} else {
		if self.head == nil {
			self.head = node
		} else {
			tempNode := self.head
			node.next = tempNode
			tempNode.prev = node
			self.head = node
		}
		self.size++
	}

}

/*插入尾部*/
func (self *LinkedList)insertTail(data interface{})  {
	node := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if self.head == nil {
		self.head = node
	} else {
		current := self.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
		if self.single == false {
			node.prev = current
		}
	}
	self.size++
}

func (self *LinkedList)DeleteByIndex(index int) bool {
	if self == nil || self.size == 0 {
		return false
	}
	if index < 0 || index >= self.size {
		return false
	}
	node := self.head
	prev := self.head
	for i := 0; i < index; i++ {
		prev = node
		if node == nil {
			return false
		}
		node = node.next
	}
	if self.single == false {
		node.next.prev = prev
		node.prev = nil
	}
	prev.next = node.next
	node.next = nil
	node.data = nil
	node = nil
	self.size--
	return true
}
/*删除指定数据节点*/
func (self *LinkedList)DeleteByData(data interface{}, all bool) bool {
	if self == nil || self.size == 0 {
		return false
	}
	node := self.head
	prev := self.head
	for node.next != nil {
		prev = node
		node = node.next
		if reflect.TypeOf(data) == reflect.TypeOf(node.data) && data == node.data {
			if !self.single {
				node.next.prev = prev
				node.prev = nil
			}
			prev.next = node.next
			node.next = nil
			self.size--
			if !all {
				return true
			}
		}
	}
	return true
}

/*根据指定下标获取节点*/
func (self *LinkedList)Get(index int) *LinkedListNode {
	if self == nil || self.size == 0 || index >= self.size {
		return nil
	}
	if index == 0 {
		return self.head
	}
	node := self.head
	for i := 1; i <= index; i++ {
		if node == nil {
			return nil
		}
		node = node.next
	}
	return node
}

func (self *LinkedList)Search(data interface{}) []int  {
	var idxs = []int{}
	node := self.head
	idx := 0
	for node.next != nil {
		if reflect.TypeOf(data) == reflect.TypeOf(node.data) && data == node.data {
			idxs = append(idxs, idx)
		}
		idx++
		node = node.next
	}
	return idxs
}

/*销毁链表*/
func (self *LinkedList)Destory()  {
	self.head.destory()
	self.size = 0
}

/*销毁链表*/
func (node *LinkedListNode)destory()  {
	if node == nil {
		return
	}
	node.next.destory()
	node.prev = nil
	node.next = nil
	node.data = nil
}




