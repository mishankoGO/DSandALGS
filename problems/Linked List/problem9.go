package main

import (
	"DSandALGS/chapter3/cll"
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Check whether the given list is either NULL-terminated or ends in a cycle */
// Time: O(n), Space: O(1)

func NullOrCyclicFloyd(linkedList ll.LinkedList) string {
	var slow, fast = linkedList.HeadNode, linkedList.HeadNode
	for {
		if fast.NextNode == nil || fast.NextNode.NextNode == nil {
			return "Null"
		} else {
			slow = slow.NextNode
			fast = fast.NextNode.NextNode
			if fast == slow {
				return "Cyclic"
			}
		}
	}
}

func CyclicOrNullFloyd(linkedList cll.CircularLinkedList) string {
	var slow, fast = linkedList.HeadNode, linkedList.HeadNode
	for {
		if fast.NextNode == nil || fast.NextNode.NextNode == nil {
			return "Null"
		} else {
			slow = slow.NextNode
			fast = fast.NextNode.NextNode
			if fast == slow {
				return "Cyclic"
			}
		}
	}
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)

	fmt.Println(NullOrCyclicFloyd(linkedList))

	clinkedList := cll.CircularLinkedList{}
	clinkedList.AddToHead(1)
	clinkedList.AddToEnd(2)
	clinkedList.AddToEnd(3)
	clinkedList.AddToEnd(4)
	clinkedList.AddToEnd(5)

	fmt.Println(CyclicOrNullFloyd(clinkedList))

}
