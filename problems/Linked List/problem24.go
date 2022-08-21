package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Find the middle of linked list */
// Time: O(n), Space: O(1)

func findMiddle(linkedList ll.LinkedList) *ll.Node {
	var middle = linkedList.HeadNode
	var length int
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		length++
	}
	fmt.Println(length / 2)
	for i := 0; i < length/2; i++ {
		middle = middle.NextNode
	}
	return middle
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

	fmt.Println(findMiddle(linkedList))

}
