package dll

import "fmt"

// Node class
type Node struct {
	Property     int
	previousNode *Node
	nextNode     *Node
}

// DoublyLinkedList class
type DoublyLinkedList struct {
	HeadNode *Node
}

// NodeBetweenValues method
func (linkedList *DoublyLinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var nodeWith *Node
	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.Property == firstProperty && node.nextNode.Property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead method
func (linkedList *DoublyLinkedList) AddToHead(property int) {
	// make a new node with desired property
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	if linkedList.HeadNode != nil {
		node.nextNode = linkedList.HeadNode     // point new node to head node
		linkedList.HeadNode.previousNode = node // point head node to new node
	}
	linkedList.HeadNode = node // make new node a head node
}

// NodeWithValue method
func (linkedList *DoublyLinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node
	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// LastNode method
func (linkedList *DoublyLinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddAfter method
func (linkedList *DoublyLinkedList) AddAfter(nodeProperty int, property int) {
	// make a new node with desired property
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	nodeWith := linkedList.NodeWithValue(nodeProperty) // find node to insert new node after
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode // point new node to next node of desired node
		node.previousNode = nodeWith      // point new node to desired node
		nodeWith.nextNode = node          // point desired node to new node
	}
}

// AddToEnd method
func (linkedList *DoublyLinkedList) AddToEnd(property int) {
	// make a new node with desired property
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	lastNode := linkedList.LastNode() // traverse to find the last node
	if lastNode != nil {
		lastNode.nextNode = node     // point last node to new node
		node.previousNode = lastNode // point new node to last node
	}
}

// IterateList method
func (linkedList *DoublyLinkedList) IterateList() {
	for node := linkedList.HeadNode; node != nil; node = node.nextNode {
		fmt.Println(node.Property)
	}
}

// Delete method
func (linkedList *DoublyLinkedList) Delete(property int) {
	if linkedList.HeadNode.Property == property {
		linkedList.HeadNode = linkedList.HeadNode.nextNode
		if linkedList.HeadNode != nil {
			linkedList.HeadNode.previousNode = nil
		}
	}

	nodeWith := linkedList.NodeWithValue(property)
	if nodeWith != nil {
		nodeWith.previousNode.nextNode = nodeWith.nextNode
		if nodeWith.nextNode != nil {
			nodeWith.nextNode.previousNode = nodeWith.previousNode
		}
	}
}
