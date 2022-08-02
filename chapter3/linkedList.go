package main

import "fmt"

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method
func (linkedList *LinkedList) AddToHead(property int) {
	var node Node
	node.property = property
	node.nextNode = linkedList.headNode
	linkedList.headNode = &node
}

//IterateList

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(5)
	fmt.Println(linkedList.headNode.property)

}
