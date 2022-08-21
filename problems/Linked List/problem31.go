package main

import "DSandALGS/chapter3/ll"

/* Merge sorted linked lists */
// Time: O(n + m), Space: O(max(n, m))

func mergeSortedLinkedLists(l1, l2 ll.LinkedList) (res ll.LinkedList) {
	var curr1, curr2 = l1.HeadNode, l2.HeadNode
	if curr1.Property > curr2.Property {
		res.AddToHead(curr2.Property)
		curr2 = curr2.NextNode
	} else {
		res.AddToHead(curr1.Property)
		curr1 = curr1.NextNode
	}
	for {
		if curr1 != nil && curr2 != nil {
			if curr1.Property > curr2.Property {
				res.AddToEnd(curr2.Property)
				curr2 = curr2.NextNode
			} else {
				res.AddToEnd(curr1.Property)
				curr1 = curr1.NextNode
			}
		} else if curr1 == nil && curr2 != nil {
			res.AddToEnd(curr2.Property)
			curr2 = curr2.NextNode
		} else if curr2 == nil && curr1 != nil {
			res.AddToEnd(curr1.Property)
			curr1 = curr1.NextNode
		} else {
			return
		}
	}
}

func main() {
	l1 := ll.LinkedList{}
	l2 := ll.LinkedList{}
	l1.AddToHead(1)
	l2.AddToHead(2)
	l1.AddToEnd(3)
	l1.AddToEnd(5)
	l1.AddToEnd(7)
	l2.AddToEnd(4)
	l2.AddToEnd(6)
	l2.AddToEnd(8)

	l3 := mergeSortedLinkedLists(l1, l2)
	l3.IterateList()
}
