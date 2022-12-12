package tasks

import (
	"fmt"
	"testing"
)

func TestSolutionE(t *testing.T) {
	var node1, node2, node3 *DoubleListNode
	node3 = &DoubleListNode{
		data: "3",
		next: nil,
		prev: node2,
	}
	node2 = &DoubleListNode{
		data: "2",
		next: node3,
		prev: node1,
	}
	node1 = &DoubleListNode{
		data: "1",
		next: node2,
		prev: nil,
	}

	// assign next and prevs in lists
	node2.prev = node1
	node2.next = node3
	node3.prev = node2
	node1.next = node2

	newHead := SolutionE(node3)
	for node := newHead; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}
