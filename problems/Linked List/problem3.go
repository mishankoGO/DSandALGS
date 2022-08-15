package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Find n-th node from the end of a Linked List using hash table */
// Time: O(n), Space: O(n)

func FindNthFromEndHashTable(linkedList *ll.LinkedList, n int) int {

	idxMap := make(map[int]*int)
	count := 0
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		idxMap[count] = &node.Property
		count++
	}
	if n > count {
		fmt.Println("Linked List is too short")
		return -1
	}
	resIdx := count - n
	return *idxMap[resIdx]
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

	res := FindNthFromEndHashTable(&linkedList, 1)
	fmt.Println(res)

}
