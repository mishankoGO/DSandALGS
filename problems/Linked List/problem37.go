package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* How will you check if the linked list is palindrome or not? */
// Time: O(n), Space: O(1)

func FindMiddleTwoPointers(linkedList ll.LinkedList) *ll.Node {
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

func palindromeLinkedList(linkedList *ll.LinkedList) bool {
	var mid *ll.Node
	var buf ll.LinkedList
	mid = FindMiddleTwoPointers(*linkedList)

	buf.HeadNode = mid
	ReverseLinkedList(&buf)

	first, second := linkedList.HeadNode, buf.HeadNode
	for first != nil && second != nil {
		if first.Property != second.Property {
			return false
		}
		first = first.NextNode
		second = second.NextNode
	}
	return true
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(1)

	fmt.Println(palindromeLinkedList(&linkedList))
}
