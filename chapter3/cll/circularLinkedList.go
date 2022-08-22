package cll

import (
	"fmt"
)

// Node class
type Node struct {
	Property int
	NextNode *Node
}

// CircularLinkedList class
type CircularLinkedList struct {
	HeadNode *Node
}

// CircularListLength method
func (linkedList *CircularLinkedList) CircularListLength() int {
	var current = linkedList.HeadNode
	var count int
	if linkedList.HeadNode == nil {
		return 0
	}
	for {
		current = current.NextNode
		count++
		if current == linkedList.HeadNode {
			break
		}
	}
	return count
}

// IterateList method
func (linkedList *CircularLinkedList) IterateList() {
	if linkedList.HeadNode == nil {
		return
	}
	var current = linkedList.HeadNode
	for {
		fmt.Println(current.Property)
		current = current.NextNode
		if current == linkedList.HeadNode {
			return
		}
	}
}

// AddToHead method
func (linkedList *CircularLinkedList) AddToHead(property int) {
	var node Node
	node.Property = property
	if linkedList.HeadNode == nil {
		node.NextNode = &node
	} else {
		lastNode := linkedList.LastNode()
		node.NextNode = linkedList.HeadNode
		lastNode.NextNode = &node
	}
	linkedList.HeadNode = &node
}

// AddToEnd method
func (linkedList *CircularLinkedList) AddToEnd(property int) {
	var node Node
	node.Property = property
	node.NextNode = linkedList.HeadNode

	lastNode := linkedList.LastNode()
	lastNode.NextNode = &node
}

// LastNode method
func (linkedList *CircularLinkedList) LastNode() *Node {
	var lastNode *Node
	if linkedList.HeadNode == nil {
		return nil
	}
	var current = linkedList.HeadNode
	for {
		current = current.NextNode
		if current.NextNode == linkedList.HeadNode {
			lastNode = current
			break
		}

	}
	return lastNode
}

// DeleteLast method
func (linkedList *CircularLinkedList) DeleteLast() {
	lastNode := linkedList.LastNode()
	previousToDelete := linkedList.HeadNode
	for {
		if previousToDelete.NextNode == lastNode {
			previousToDelete.NextNode = linkedList.HeadNode
			lastNode.NextNode = lastNode
			return
		}
		previousToDelete = previousToDelete.NextNode
	}
}

// DeleteHead method
func (linkedList *CircularLinkedList) DeleteHead() {
	//var auxNode = linkedList.HeadNode
	lastNode := linkedList.LastNode()
	lastNode.NextNode = linkedList.HeadNode.NextNode
	linkedList.HeadNode = linkedList.HeadNode.NextNode
}
