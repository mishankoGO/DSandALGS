package linkedLists

type Node struct {
	Data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n

	l.head.next = second
	l.length++
}
