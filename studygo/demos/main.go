package main

import (
	"fmt"
	"github.com/znk_fullstack/studygo/demos/chain"
)

func main() {
	list := chain.CreateLinkedList(true)
	for i := 0; i < 12; i++ {
		list.InsertLinkedListNode(-1, i)
	}
	head := list.Head()
	//head.Print(true, true)
	//fmt.Println("linked list length", list.Length())

	fmt.Println("+++++++++++0++++++++++")

	list.InsertLinkedListNode(5, 20)
	
	head.Print(true, true)
	//fmt.Println("linked list length", list.Length())
	
	node := list.Get(5)
	node.Print(false, true)
	
	fmt.Println("+++++++++++++1+++++++++++++")
	
	succ := list.DeleteByIndex(5)
	fmt.Println("delete succ: ", succ)
	head.Print(true, true)
	
	fmt.Println("+++++++++++++2+++++++++++++")
	list.Destory()
	head.Print(true, true)
}
