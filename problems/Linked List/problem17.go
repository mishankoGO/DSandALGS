package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Give an algorithm for finding the merging point */
// Time: O(n), Space: O(min(n, m))

func mergingPoint(l1, l2 ll.LinkedList) *ll.Node {
	pointerMap := make(map[*ll.Node]bool)
	for node := l1.HeadNode; node != nil; node = node.NextNode {
		pointerMap[node] = true
	}
	for node := l2.HeadNode; node != nil; node = node.NextNode {
		if _, ok := pointerMap[node]; ok {
			return node
		}
	}
	return nil
}

func main() {
	linkedList1 := ll.LinkedList{}
	linkedList1.AddToHead(1)
	linkedList1.AddToEnd(2)
	linkedList1.AddToEnd(3)
	linkedList1.AddToEnd(4)
	linkedList1.AddToEnd(5)
	linkedList1.AddToEnd(6)
	linkedList1.AddToEnd(7)

	linkedList2 := ll.LinkedList{}
	linkedList2.AddToHead(1)
	linkedList2.AddToEnd(2)
	//linkedList2.LastNode().NextNode = linkedList1.NodeWithValue(6)

	linkedList1.IterateList()
	fmt.Println("next list")
	linkedList2.IterateList()

	fmt.Println(mergingPoint(linkedList1, linkedList2))
}
