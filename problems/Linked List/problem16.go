package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Reverse a singly linked list. */
// Time: O(n), Space: O(1)

func ReverseLinkedList(linkedList *ll.LinkedList) {
	var prev, next *ll.Node

	for curr := linkedList.HeadNode; curr != nil; {
		next = curr.NextNode // save next node
		curr.NextNode = prev // point to the previous node
		prev = curr          // point previous to it's new next
		curr = next          // curr = curr.NextNode
	}
	linkedList.HeadNode = prev
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(4)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(1)
	linkedList.IterateList()
	fmt.Print("\n")
	ReverseLinkedList(&linkedList)
	linkedList.IterateList()
}
