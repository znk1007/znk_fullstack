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
	//fmt.Println("-------insert------ 1")
	that := homework03.CreateLinkedList()
	//for i := 0; i < 10; i++ {
	//	that.Insert(i, -1)
	//}
	//that.Print()
	//
	//fmt.Println("--------1------- length: ", that.Length())
	//
	//fmt.Println("-------insert------ 2")
	//
	//for i := 10; i < 20; i++ {
	//	that.Insert(i, that.Length()+1)
	//}
	//that.Print()
	//
	//fmt.Println("--------2------- length: ", that.Length())
	//
	//
	//fmt.Println("--------2------- is circular: ", that.IsCircular())
	//
	//fmt.Println("--------3------- length: ", that.Length())
	//fmt.Println("tail: ", that.Tail())
	//node := that.Get(10)
	//fmt.Println("get node: ", node)
	//
	//fmt.Println("--------4-------")
	//for i := 20; i < 30; i++ {
	//	that.Insert(i, 2)
	//}
	//
	//that.Print()
	
	fmt.Println("--------5-------")

	for i := 1; i < 11; i++ {
		that.Insert(i, that.Length()+1)
	}

	that.Print()
	fmt.Println("--------5------- is circular: ", that.IsCircular())
	fmt.Println("--------5------- length: ", that.Length())

	fmt.Println("--------6------- delete: ")

	that.DeleteByIndex(2)

	that.Print()
	fmt.Println("--------6------- is circular: ", that.IsCircular())
	fmt.Println("--------6------- length: ", that.Length())

	fmt.Println("--------7-------")

	node := that.Search(10, true)
	fmt.Println("search node: ", node)

	fmt.Println("--------8------- delete")

	that.DeleteByData(10, true)
	that.Print()
	fmt.Println("--------8------- head: ", that.Head())
	fmt.Println("--------8------- tail: ", that.Tail())
	fmt.Println("--------8------- is circular: ", that.IsCircular())
	fmt.Println("--------8------- length: ", that.Length())
	
}
