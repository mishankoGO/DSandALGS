package linkedLists

import "fmt"

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

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *LinkedList) Delete(i int) {
	if l.length == 0 {
		return
	}

	if l.head.Data == i {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.Data != i {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = l.head.next
	}
	// if next node data == i -> skip one node
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
