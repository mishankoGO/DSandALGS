package tasks

import "fmt"

func SolutionC(head *ListNode, index int) *ListNode {
	var newHead *ListNode
	if index == 0 {
		newHead = head.next
		head.next = nil
		return newHead
	}
	for node := head; index != 0; index-- {
		if index == 1 {
			nextNode := node.next
			if node.next == nil && index > 0 {
				fmt.Println("index is out of range")
				return nil
			}
			node.next = nextNode.next
			nextNode.next = nil
		}
		node = node.next
	}

	return head
}
