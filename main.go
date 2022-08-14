package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

func main() {
	var a ll.LinkedList
	a.AddToHead(1)
	a.AddToEnd(2)
	a.AddToEnd(3)

	fmt.Println("before deletion")
	a.IterateList()

	a.Delete(3)

	fmt.Println("after deletion")
	a.IterateList()

	a.DeleteLinkedList()

	fmt.Println("after full deletion")
	fmt.Println(a)

}
