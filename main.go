package main

import (
	"DSandALGS/chapter3/dll"
	"fmt"
)

func main() {
	var a dll.DoublyLinkedList
	a.AddToHead(1)
	a.AddToEnd(2)
	a.AddToEnd(4)
	a.AddAfter(2, 3)

	fmt.Println("before deletion")
	a.IterateList()

	a.Delete(3)

	fmt.Println("after deletion")
	a.IterateList()
}
