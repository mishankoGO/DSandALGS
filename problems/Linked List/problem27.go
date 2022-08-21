package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Find the middle of linked list (Two Pointers) */
// Time: O(n), Space: O(1)

func findMiddleTwoPointers(linkedList ll.LinkedList) *ll.Node {
	var p1, p2 = linkedList.HeadNode, linkedList.HeadNode
	for p2 != nil {
		if p2.NextNode != nil {
			p2 = p2.NextNode.NextNode
		} else {
			break
		}
		p1 = p1.NextNode
	}

	return p1
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)
	linkedList.AddToEnd(6)
	linkedList.AddToEnd(7)
	linkedList.AddToEnd(8)
	linkedList.AddToEnd(9)

	fmt.Println(findMiddleTwoPointers(linkedList))

}
