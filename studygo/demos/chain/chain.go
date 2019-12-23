package chain
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

func (list *LinkedList)insert(byHead bool, datas ...interface{}) bool {
	dataLen := len(datas)
	if dataLen == 0 {
		 return false
	}
	node := list.head
	for i := 0; i < dataLen; i++ {
		newNode := &LinkedListNode{
			data: datas[i],
			prev: nil,
			next: nil,
		}
		if node == nil {
			list.head = newNode
			node = list.head
		}
		
	}
}

/*插入节点*/
func (list *LinkedList)InsertLinkedListNode(index int, byHead bool, datas ...interface{}) bool  {
	if list == nil || datas == nil || len(datas) == 0 || list.size < index {
		return false
	}
	if list.single {
		if list.size == 0 {
			node := list.head
			for i := 0; i < len(datas); i++ {
				newNode := &LinkedListNode{
					data: datas[i],
				}
				if list.head == nil {
					list.head = newNode
					node = newNode
				}
				if byHead {
					newNode.next = node.next
					node.next = newNode
				} else {
					node.next = newNode
					node = newNode
				}
				list.size++
			}
		} else {
			node := list.head
			if index < 1 || index >= list.size {
				newNode := &LinkedListNode{
					data: data,
					prev: nil,
					next: nil,
				}
				if byHead {
					newNode.next = node.next
					node.next = newNode
				} else {
					for node.next != nil {
						node = node.next
					}
					node.next = newNode
				}
				list.size++
				return true
			}
			i := 1
			for i < index || node != nil  {
				node = node.next
				i++
			}
			if i > index || node == nil {
				return false
			}
			newNode := &LinkedListNode{
				data: data,
				prev: nil,
				next: nil,
			}
			newNode.next = node.next
			node.next = newNode
			list.size++
			return true
		}
	} else {
		
	}
	
	return false
}





