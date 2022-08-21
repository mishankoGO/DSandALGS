package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Give an algorithm for finding the merging point (Lengths) */
// Time: O(max(n, m)), Space: O(1)

func mergingPointLengths(l1, l2 ll.LinkedList) *ll.Node {
	var length1, length2 int
	var node1, node2 = l1.HeadNode, l2.HeadNode
	for {
		if node1 != nil {
			length1++
			node1 = node1.NextNode
		}
		if node2 != nil {
			length2++
			node2 = node2.NextNode
		}
		if node2 == nil && node1 == nil {
			break
		}
	}

	if length1 > length2 {
		d := length1 - length2
		node1, node2 = l1.HeadNode, l2.HeadNode
		for i := 0; i < d; i++ {
			node1 = node1.NextNode
		}
		for node1 != node2 {
			node1 = node1.NextNode
			node2 = node2.NextNode
		}
		return node1
	} else {
		d := length2 - length1
		node1, node2 = l1.HeadNode, l2.HeadNode
		for i := 0; i < d; i++ {
			node1 = node1.NextNode
		}
		for node1 != node2 {
			node1 = node1.NextNode
			node2 = node2.NextNode
		}
		return node1
	}
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
	//linkedList2.LastNode().NextNode = linkedList1.NodeWithValue()

	linkedList1.IterateList()
	fmt.Println("next list")
	linkedList2.IterateList()

	fmt.Println(mergingPointLengths(linkedList1, linkedList2))
}
