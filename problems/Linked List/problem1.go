package main

import (
	"DSandALGS/chapter3/ll"
	"fmt"
)

/* Implement Stack Using Linked List */

// Stack class
type Stack struct {
	items ll.LinkedList
}

// Push method
func (s *Stack) Push(i int) {
	s.items.AddToHead(i)
}

// Peek method
func (s *Stack) Peek() int {
	return s.items.HeadNode.Property
}

// Pop method
func (s *Stack) Pop() int {
	if s.items.HeadNode == nil {
		fmt.Println("No elements in a Stack")
		return -1
	}
	res := s.items.HeadNode.Property          // last added element
	s.items.Delete(s.items.HeadNode.Property) // delete it from list
	return res
}

func main() {
	stack := Stack{}
	stack.Push(4)
	stack.Peek()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	stack.items.IterateList()

	a := stack.Pop()
	fmt.Print("popped elem: ")
	fmt.Println(a)

	fmt.Println("all list")
	stack.items.IterateList()
}
