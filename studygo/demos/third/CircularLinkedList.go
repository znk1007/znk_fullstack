package linkedList

type LinkedList struct {
	 single bool
	 size int
	 head *LinkedListNode
	 tail *LinkedListNode
}

type LinkedListNode struct {
	data interface{}
	prev *LinkedListNode
	next *LinkedListNode
}