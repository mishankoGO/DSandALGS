package dll

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
	var node = &Node{}
	node.Property = property
	node.nextNode = nil
	if linkedList.HeadNode != nil {
		node.nextNode = linkedList.HeadNode
		linkedList.HeadNode.previousNode = node
	}
	linkedList.HeadNode = node
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
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// AddToEnd method
func (linkedList *DoublyLinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.Property = property
	node.nextNode = nil

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}
