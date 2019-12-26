package main

import (
	"fmt"
	homework03 "github.com/znk_fullstack/studygo/demos/third"
)

func main() {
	//list := chain.CreateLinkedList(true)
	//for i := 0; i < 12; i++ {
	//	list.InsertLinkedListNode(-1, i)
	//}
	////head := list.Head()
	////head.Print(true, true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++0++++++++++")
	//
	//list.InsertLinkedListNode(5, 20)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//node := list.Get(5)
	//fmt.Println("the node: ", node)
	//
	//fmt.Println("+++++++++++++1+++++++++++++")
	//
	//list.InsertLinkedListNode(8, 20)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++++2+++++++++++++")
	//idxs := list.Search(20)
	//fmt.Println("idxs: ", idxs)
	//
	//fmt.Println("+++++++++++++3+++++++++++++")
	//
	//succ := list.DeleteByIndex(5)
	//fmt.Println("delete succ: ", succ)
	//list.Print(true)
	//fmt.Println("linked list length", list.Length())
	//
	//fmt.Println("+++++++++++++4+++++++++++++")
	//list.Destory()
	//list.Print(true)
	fmt.Println("-------insert------ 1")
	that := homework03.CreateLinkedList()
	for i := 0; i < 10; i++ {
		that.Insert(i, -1)
	}
	that.Print()

	fmt.Println("--------1------- length: ", that.Length())

	fmt.Println("-------insert------ 2")

	for i := 11; i < 20; i++ {
		that.Insert(i, that.Length()+1)
	}
	that.Print()

	fmt.Println("--------2------- length: ", that.Length())


	fmt.Println("is circular: ", that.IsCircular())
}
