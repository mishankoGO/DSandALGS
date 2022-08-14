package main

import (
	"DSandALGS/chapter3/cll"
	"fmt"
)

func main() {
	var a cll.CircularLinkedList
	a.AddToHead(2)
	a.AddToHead(3)
	a.AddToHead(4)
	a.AddToHead(5)
	fmt.Println(a.CircularListLength())
	a.IterateList()
	a.AddToEnd(1)
	fmt.Println(a.CircularListLength())
	a.IterateList()
	a.DeleteLast()
	fmt.Println("after deletion")
	a.IterateList()
	a.DeleteHead()
	fmt.Println("after deletion")
	a.IterateList()
}
