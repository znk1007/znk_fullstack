package chain

import "fmt"

/*链表对象*/
type LinkedList struct {
	size int
	single bool
	close bool
	head *LinkedListNode
}

/*链表节点对象*/
type LinkedListNode struct {
	 data interface{}
	 prev *LinkedListNode
	 next *LinkedListNode
}

/*创建链表*/
func CreateLinkedList(single bool, close bool) *LinkedList  {
	linkedList := &LinkedList{
		size:   0,
		single: single,
		close:close,
		head:   nil,
	}
	return linkedList
}

/*插入节点*/
func (list *LinkedList)InsertLinkedListNode(index int, data interface{}) bool  {
	if list == nil || data == nil {
		return false
	}
	if index < 0 {
		list.insertHead(data)
	} else if index > list.size {
		list.insertTail(data)
	} else {
		node := list.head
		for i := 0; i < index - 1; i++ {
			if node == nil {
				return false
			}
			node = node.next
		}
		newNode := &LinkedListNode{
			data: data,
			prev: nil,
			next: nil,
		}
		newNode.next = node.next
		node.next = newNode
		if list.single == false {
			newNode.prev = node
		}
		list.size++
	}
	return false
}
/*链表长度*/
func (list *LinkedList) Length() int {
	return list.size
}

/*打印链表数据*/
func (node *LinkedListNode)Print(loop bool)  {
	if node == nil {
		return
	}
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
	if loop {
		node.next.Print(loop)
	}
}

/*头部节点*/
func (list *LinkedList) Head() *LinkedListNode  {
	return list.head
}

func (list *LinkedList)Get(index int) *LinkedListNode {
	if list == nil || list.size == 0 || index >= list.size {
		return nil
	}
	if index == 0 {
		 return list.head
	}
	node := list.head
	for i := 1; i <= index; i++ {
		if node == nil {
			 return nil
		}
		node = node.next
	}
	return node
}

/*插入头部*/
func (list *LinkedList)insertHead(data interface{})  {
	if list == nil || data == nil {
		return
	}
	node := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if list.single {
		node.next = list.head
		list.head = node
		list.size++
	} else {
		node.next = list.head
		if list.head != nil {
			node.prev = list.head.prev
		}
		list.head = node
		list.size++
	}

}

/*插入尾部*/
func (list *LinkedList)insertTail(data interface{})  {
	node := &LinkedListNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
		node.prev = current
	}
	list.size++
}






