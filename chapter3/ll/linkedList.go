package ll

import "fmt"

// Node class
type Node struct {
	Property int
	NextNode *Node
}

// LinkedList class
type LinkedList struct {
	HeadNode *Node
}

// AddToHead method
func (linkedList *LinkedList) AddToHead(property int) {
	var node Node
	node.Property = property
	node.NextNode = linkedList.HeadNode
	linkedList.HeadNode = &node
}

// IterateList method
func (linkedList *LinkedList) IterateList() {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.Property)
	}
}

// LastNode method
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.Property = property
	node.NextNode = nil
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = node
	}
}

// NodeWithValue method
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.Property = property
	node.NextNode = nil

	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}
