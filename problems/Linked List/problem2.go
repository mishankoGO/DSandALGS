package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Find n-th node from the end of a Linked List */

func FindNthFromEnd(linkedList *ll.LinkedList, n int) int {
	var count int
	var res int
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		count++
	}
	fmt.Println(count)
	if n > count {
		fmt.Println("Linked List is too short")
		return -1
	} else {
		resIdx := count - n
		var i int
		for node := linkedList.HeadNode; i <= resIdx; node = node.NextNode {
			if i == resIdx {
				res = node.Property
			}
			i++
		}
	}
	return res
}

func main() {
	linkedList := ll.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)

	linkedList.IterateList()
	fmt.Println()

	res := FindNthFromEnd(&linkedList, 6)
	fmt.Println(res)

}
