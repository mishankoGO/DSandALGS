package main

import (
	"DSandALGS/chapter3/cll"
	"fmt"
)

/* Split a Circular Linked List into two equal parts. */
// Time: O(n), Space: O(n)

func CyclicFloyd(linkedList *cll.CircularLinkedList, middle int) (*cll.Node, *cll.Node) {
	var slow, fast, middleNode = linkedList.HeadNode, linkedList.HeadNode, linkedList.HeadNode
	for {
		middle--
		if middle == 0 {
			middleNode = slow
		}
		if fast.NextNode == nil || fast.NextNode.NextNode == nil {
			return nil, nil
		} else {
			if slow.NextNode == fast.NextNode.NextNode {
				return slow, middleNode
			} else {
				slow = slow.NextNode
				fast = fast.NextNode.NextNode
			}
		}
	}
}

func splitCircularLL(ll *cll.CircularLinkedList) (left, right cll.CircularLinkedList) {
	length := ll.CircularListLength()

	middlePoint := length / 2
	if 2*middlePoint < length {
		middlePoint = middlePoint + 1
	}
	last, middle := CyclicFloyd(ll, middlePoint)
	left.HeadNode = ll.HeadNode
	right.HeadNode = middle.NextNode

	middle.NextNode = left.HeadNode
	last.NextNode = right.HeadNode

	return left, right
}

func main() {
	ll := cll.CircularLinkedList{}
	ll.AddToHead(1)
	ll.AddToEnd(2)
	ll.AddToEnd(3)
	ll.AddToEnd(4)
	ll.AddToEnd(5)
	ll.AddToEnd(6)
	ll.AddToEnd(7)

	ll.IterateList()
	fmt.Println()
	left, right := splitCircularLL(&ll)
	left.IterateList()
	fmt.Println()
	right.IterateList()
}
