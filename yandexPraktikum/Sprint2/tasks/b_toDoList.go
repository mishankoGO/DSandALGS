package tasks

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func SolutionB(head *ListNode) {
	for node := head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}
