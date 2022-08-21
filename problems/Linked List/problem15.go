package main

import "DSandALGS/chapter3/ll"

/* Insert a node in a sorted linked list */
// Time: O(n), Space: O(1)

func InsertIntoSortedList(linkedList *ll.LinkedList, i int) {
	for node := linkedList.HeadNode; node.NextNode != nil; node = node.NextNode {
		if linkedList.HeadNode.Property >= i {
			linkedList.AddToHead(i)
			return
		}
		if i <= node.NextNode.Property {
			linkedList.AddAfter(node.Property, i)
			return
		}
	}
	linkedList.AddToEnd(i)
}

func main() {
	sortedLinkedList := ll.LinkedList{}
	sortedLinkedList.AddToHead(1)
	sortedLinkedList.AddToEnd(2)
	sortedLinkedList.AddToEnd(4)
	sortedLinkedList.AddToEnd(5)

	InsertIntoSortedList(&sortedLinkedList, 0)

	sortedLinkedList.IterateList()
}
