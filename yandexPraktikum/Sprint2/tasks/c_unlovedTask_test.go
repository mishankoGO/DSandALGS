package tasks

import (
	"fmt"
	"testing"
)

func TestSolutionC(t *testing.T) {
	node3 := &ListNode{
		data: "3",
		next: nil,
	}
	node2 := &ListNode{
		data: "2",
		next: node3,
	}
	nodeToDelete := &ListNode{
		data: "To delete",
		next: node2,
	}
	node1 := &ListNode{
		data: "1",
		next: nodeToDelete,
	}
	SolutionB(node1)

	newNode := SolutionC(node1, 4)
	fmt.Println("------------------")

	SolutionB(newNode)
}
