package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Reverse blocks of K nodes in a list. */
// Time: O(N), Space: O(1)

/* The following steps are needed to arrive at the desired output.
1) Create a dummy node. Point next of this node to head of the linked list provided.
2) Iterate through the given linked list to get the length of the list.
3) Create three pointer prev,curr and next to reverse each group. Why? If we observe output,
	we can see that we have to reverse each group, apart from modifying links of group.
	Init prev with dummy
4) Iterate through the linked list until the length of list to be processed is greater than and equal to given k.
5) For each iteration, point cur to pre->next and nex to curr->next.
6) Start nested iteration for length of k.
7) For each iteration, modify links as following steps:-
	curr->next = next->next
	next->next = prev->next
	prev->next = next
	next = curr->next
	Move prev to curr and reduce length by k.
*/

func lengthLinkedList(linkedList *ll.LinkedList) int {
	var length int
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		length++
	}
	return length
}

func reverseKnodes(linkedList *ll.LinkedList, k int) {
	if linkedList.HeadNode == nil || k <= 1 || linkedList.HeadNode.NextNode == nil {
		return
	}

	var dummy = &ll.Node{Property: -1}
	dummy.NextNode = linkedList.HeadNode

	length := lengthLinkedList(linkedList)
	var prev, curr, next *ll.Node
	prev = dummy

	for length >= k {
		curr = prev.NextNode
		next = curr.NextNode

		for i := 1; i < k; i++ {
			curr.NextNode = next.NextNode
			next.NextNode = prev.NextNode
			prev.NextNode = next
			next = curr.NextNode
		}
		prev = curr
		length -= k
	}
	linkedList.HeadNode = dummy.NextNode
}

//func lengthOfLinkedList(linkedList *ll.LinkedList) int {
//	var length int
//	for curr := linkedList.HeadNode; curr != nil; curr = curr.NextNode {
//		length++
//	}
//	return length
//}
//
//func reverseKnodes(linkedList *ll.LinkedList, k int) {
//	// if head is nil or k is 1 then return head
//	if linkedList.HeadNode == nil || k <= 1 || linkedList.HeadNode.NextNode == nil {
//		return
//	}
//
//	var length = lengthOfLinkedList(linkedList)
//
//	var dummy = &ll.Node{Property: -1}
//	dummy.NextNode = linkedList.HeadNode
//
//	var prev, curr, next *ll.Node
//	prev = dummy
//
//	for length >= k {
//		curr = prev.NextNode
//		next = curr.NextNode
//		for i := 1; i < k; i++ {
//			curr.NextNode = next.NextNode
//			next.NextNode = prev.NextNode
//			prev.NextNode = next
//			next = curr.NextNode
//		}
//		prev = curr
//		length -= k
//	}
//	linkedList.HeadNode = dummy.NextNode
//}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(7)
	linkedList.AddToEnd(6)
	linkedList.AddToEnd(5)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(1)
	linkedList.IterateList()
	fmt.Print("\n")
	reverseKnodes(&linkedList, 3)
	linkedList.IterateList()
}
