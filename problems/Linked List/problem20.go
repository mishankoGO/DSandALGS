package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Give an algorithm for finding the merging point (Stacks) */
// Time: O(n + m), Space: O(n + m)

// StackNode class
type StackNode struct {
	Items []*ll.Node
}

// Push method
func (s *StackNode) Push(i *ll.Node) {
	s.Items = append(s.Items, i)
}

// Pop method
func (s *StackNode) Pop() *ll.Node {
	l := len(s.Items) - 1
	item := s.Items[l]
	s.Items = s.Items[:l]
	return item
}

// Peek method
func (s *StackNode) Peek() *ll.Node {
	l := len(s.Items) - 1
	item := s.Items[l]
	return item
}

func mergingPointStack(l1, l2 ll.LinkedList) (point *ll.Node) {
	s1 := StackNode{}
	s2 := StackNode{}
	for node := l1.HeadNode; node != nil; node = node.NextNode {
		s1.Push(node)
	}
	for node := l2.HeadNode; node != nil; node = node.NextNode {
		s2.Push(node)
	}

	for len(s1.Items) > 0 {
		elem1, elem2 := s1.Peek(), s2.Peek()
		if elem1 == elem2 {
			point = s1.Pop()
			s2.Pop()
		} else {
			return point
		}
	}
	return
}

func main() {
	linkedList1 := ll.LinkedList{}
	linkedList1.AddToHead(1)
	linkedList1.AddToEnd(2)
	linkedList1.AddToEnd(3)
	linkedList1.AddToEnd(4)
	linkedList1.AddToEnd(5)
	linkedList1.AddToEnd(6)
	linkedList1.AddToEnd(7)

	linkedList2 := ll.LinkedList{}
	linkedList2.AddToHead(1)
	linkedList2.AddToEnd(2)
	//linkedList2.LastNode().NextNode = linkedList1.NodeWithValue()

	linkedList1.IterateList()
	fmt.Println("next list")
	linkedList2.IterateList()

	fmt.Println(mergingPointStack(linkedList1, linkedList2))
}
