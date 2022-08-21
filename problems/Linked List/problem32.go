package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Check whether the given Linked List length is even or odd */
// Time: O(n), Space: O(1)

func OddOrEven(linkedLink ll.LinkedList) string {
	var p1, p2 = linkedLink.HeadNode, linkedLink.HeadNode
	for p2 != nil {
		if p2.NextNode == nil {
			return "Odd"
		}
		p1 = p1.NextNode
		p2 = p2.NextNode.NextNode
	}
	return "Even"
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)

	fmt.Println(OddOrEven(linkedList))
}
