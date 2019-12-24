package main

import (
	"fmt"
	"github.com/znk_fullstack/studygo/demos/chain"
)

func main() {
	list := chain.CreateLinkedList(true,false)
	for i := 0; i < 12; i++ {
		list.InsertLinkedListNode(-1, i)
	}
	head := list.Head()
	head.Print(true)
	fmt.Println("linked list length", list.Length())

	fmt.Println("+++++++++++++++++++++")

	list.InsertLinkedListNode(5, 20)

	head.Print(true)
	fmt.Println("linked list length", list.Length())

	node := list.Get(5)
	node.Print(false)

}
