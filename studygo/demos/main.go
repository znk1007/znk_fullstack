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
	//head := list.Head()
	//head.Print(true, true)
	fmt.Println("linked list length", list.Length())

	fmt.Println("+++++++++++0++++++++++")

	list.InsertLinkedListNode(5, 20)
	list.Print(true)
	fmt.Println("linked list length", list.Length())
	
	node := list.Get(5)
	fmt.Println("the node: ", node)
	
	fmt.Println("+++++++++++++1+++++++++++++")
	
	list.InsertLinkedListNode(8, 20)
	list.Print(true)
	fmt.Println("linked list length", list.Length())
	
	fmt.Println("+++++++++++++2+++++++++++++")
	idxs := list.Search(20)
	fmt.Println("idxs: ", idxs)
	
	fmt.Println("+++++++++++++3+++++++++++++")
	
	succ := list.DeleteByIndex(5)
	fmt.Println("delete succ: ", succ)
	list.Print(true)
	fmt.Println("linked list length", list.Length())
	
	fmt.Println("+++++++++++++4+++++++++++++")
	list.Destory()
	list.Print(true)
	
	
	
	
}
