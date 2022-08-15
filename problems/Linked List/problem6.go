package main

import (
	"DSandALGS/chapter3/cll"
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Check whether the given list is either NULL-terminated or ends in a cycle */
// Time: O(n), Space: O(1)

func NullOrCyclic(linkedList ll.LinkedList) string {
	pointersMap := make(map[*ll.Node]bool)
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			return "Null"
		} else if _, ok := pointersMap[node]; ok {
			return "Cyclic"
		}
		pointersMap[node] = true

	}
	return "Unknown"
}

func CyclicOrNull(linkedList cll.CircularLinkedList) string {
	pointersMap := make(map[*cll.Node]bool)
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			return "Null"
		} else if _, ok := pointersMap[node]; ok {
			return "Cyclic"
		}
		pointersMap[node] = true

	}
	return "Unknown"
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)

	fmt.Println(NullOrCyclic(linkedList))

	clinkedList := cll.CircularLinkedList{}
	clinkedList.AddToHead(1)
	clinkedList.AddToEnd(2)
	clinkedList.AddToEnd(3)
	clinkedList.AddToEnd(4)
	clinkedList.AddToEnd(5)

	fmt.Println(CyclicOrNull(clinkedList))

}
