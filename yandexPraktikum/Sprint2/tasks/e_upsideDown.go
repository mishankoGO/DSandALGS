package tasks

type DoubleListNode struct {
	data string
	next *DoubleListNode
	prev *DoubleListNode
}

func SolutionE(head *DoubleListNode) *DoubleListNode {
	var newHead *DoubleListNode
	for node := head; node != nil; node = node.next {
		if node.next == nil { // reached the end of the linked list
			newHead = node
			break
		}
	}

	// traverse the list in backward manner and change next and prev links
	for node := newHead; node != nil; node = node.prev {
		node.next = node.prev
		node.prev = node.next
	}
	return newHead
}
