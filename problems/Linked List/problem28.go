package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Display linked list in reversed order */
/* Time: O(n), Space: O(1) */

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)

	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		defer fmt.Println(node.Property)
	}
}
