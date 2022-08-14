package cll

import (
	"fmt"
)

// Node class
type Node struct {
	property int
	nextNode *Node
}

// CircularLinkedList class
type CircularLinkedList struct {
	HeadNode *Node
}

// CircularListLength method
func (linkedLists *CircularLinkedList) CircularListLength() int {
	var current = linkedLists.HeadNode
	var count int
	if linkedLists.HeadNode == nil {
		return 0
	}
	for {
		current = current.nextNode
		count++
		if current == linkedLists.HeadNode {
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
		fmt.Println(current)
		current = current.nextNode
		if current == linkedList.HeadNode {
			return
		}
	}
}

// AddToHead method
func (linkedList *CircularLinkedList) AddToHead(property int) {
	var node Node
	node.property = property
	if linkedList.HeadNode == nil {
		node.nextNode = &node
	} else {
		lastNode := linkedList.LastNode()
		node.nextNode = linkedList.HeadNode
		lastNode.nextNode = &node
	}
	linkedList.HeadNode = &node
}

// AddToEnd method
func (linkedList *CircularLinkedList) AddToEnd(property int) {
	var node Node
	node.property = property
	node.nextNode = linkedList.HeadNode

	lastNode := linkedList.LastNode()
	lastNode.nextNode = &node
}

// LastNode method
func (linkedList *CircularLinkedList) LastNode() *Node {
	var lastNode *Node
	if linkedList.HeadNode == nil {
		return nil
	}
	var current = linkedList.HeadNode
	for {
		current = current.nextNode
		if current.nextNode == linkedList.HeadNode {
			lastNode = current
			break
		}

	}
	return lastNode
}

// DeleteLast method
func (linkedLists *CircularLinkedList) DeleteLast() {
	lastNode := linkedLists.LastNode()
	previousToDelete := linkedLists.HeadNode
	for {
		if previousToDelete.nextNode == lastNode {
			previousToDelete.nextNode = linkedLists.HeadNode
			lastNode.nextNode = lastNode
			return
		}
		previousToDelete = previousToDelete.nextNode
	}
}

// DeleteHead method
func (linkedList *CircularLinkedList) DeleteHead() {
	//var auxNode = linkedList.HeadNode
	lastNode := linkedList.LastNode()
	lastNode.nextNode = linkedList.HeadNode.nextNode
	linkedList.HeadNode = linkedList.HeadNode.nextNode
}
