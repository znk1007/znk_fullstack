package chain
/*链表对象*/
type LinkedList struct {
	size int
	single bool
	head *LinkedListNode
	tail *LinkedListNode
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
		tail:   nil,
	}
	return linkedList
}

/*插入节点*/
func (list *LinkedList)InsertLinkedListNode(data interface{}, index int, byHead bool) bool  {
	if list == nil || data == nil || list.size < index {
		return false
	}
	if list.single {
		if list.size == 0 {
			list.head = &LinkedListNode{
				data: data,
			}
			list.size++
		} else {
			tempNode := list.head
			if index < 1 || index > list.size {
				for tempNode != nil {
					tempNode = tempNode.next
				}

			}
			i := 1
			for i < index || tempNode != nil  {
				tempNode = tempNode.next
				i++
			}
			if i > index || tempNode == nil {
				return false
			}
			node := &LinkedListNode{
				data: data,
				prev: nil,
				next: nil,
			}
			node.next = tempNode.next
			tempNode.next = node
			list.size++
		}
	} else {
		
	}
	
	return false
}





