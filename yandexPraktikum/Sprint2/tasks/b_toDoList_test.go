package tasks

import "testing"

func TestSolutionB(t *testing.T) {
	node3 := &ListNode{
		data: "3",
		next: nil,
	}
	node2 := &ListNode{
		data: "2",
		next: node3,
	}
	node1 := &ListNode{
		data: "1",
		next: node2,
	}
	SolutionB(node1)
}
