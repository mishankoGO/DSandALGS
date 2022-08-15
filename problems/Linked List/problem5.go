package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Find n-th node from the end of a Linked List using two pointers */
// Time: O(n), Space: O(1)

func FindNthFromEndOneScan(linkedList *ll.LinkedList, n int) int {

	var left = linkedList.HeadNode
	count := 0
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		count++
		if count > n {
			left = left.NextNode
		}
	}
	if n > count {
		fmt.Println("Linked List is too short")
		return -1
	}

	return left.Property
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)

	linkedList.IterateList()
	fmt.Println()

	res := FindNthFromEndOneScan(&linkedList, 6)
	fmt.Println(res)

}
