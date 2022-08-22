package main

import "DSandALGS/chapter3/ll"

// TODO Understand the solution
/* Reverse the linked list in pairs */
// Time: O(n), Space: O(1)

func reverseInPairs(l *ll.LinkedList) {
	var next, prev *ll.Node

	for curr := l.HeadNode; curr != nil && curr.NextNode != nil; {
		if next != nil {
			next.NextNode.NextNode = curr.NextNode
		}
		next = curr.NextNode                   // save next pointer to the next node
		curr.NextNode = curr.NextNode.NextNode // point next node to the node after it
		next.NextNode = curr                   //
		if prev == nil {
			prev = next
		}
		curr = curr.NextNode

	}
	l.HeadNode = prev
}

func main() {
	l1 := ll.LinkedList{}
	l1.AddToHead(1)
	l1.AddToEnd(2)
	l1.AddToEnd(3)
	l1.AddToEnd(4)

	reverseInPairs(&l1)
	l1.IterateList()
}
