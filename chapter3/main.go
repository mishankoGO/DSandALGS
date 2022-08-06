package main

import (
	"DSandALGS/JunminLeeCourse/DataStructures/linkedLists"
	"fmt"
)

func main() {
	mylist := linkedLists.LinkedList{}
	node1 := &linkedLists.Node{Data: 48}
	node2 := &linkedLists.Node{Data: 18}
	node3 := &linkedLists.Node{Data: 68}

	mylist.Prepend(node1)
	mylist.Prepend(node2)
	mylist.Prepend(node3)

	fmt.Println(mylist)
}
