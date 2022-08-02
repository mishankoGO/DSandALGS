package main

import (
	"DSandALGS/chapter3/dll"
	"fmt"
)

func main() {

	linkedList := dll.DoublyLinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.HeadNode.Property)

	node := linkedList.NodeBetweenValues(1, 5)

	fmt.Println(node.Property)
}
