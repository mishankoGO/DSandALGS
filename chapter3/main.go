package main

import (
	"DSandALGS/JunminLeeCourse/DataStructures/binarySearchTrees"
	"fmt"
)

func main() {
	tree := &binarySearchTrees.Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(34)
	tree.Insert(76)
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(90)

	fmt.Println(tree.Search(200))
	fmt.Println(binarySearchTrees.Count)
}
