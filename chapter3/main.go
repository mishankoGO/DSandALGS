package main

import (
	"DSandALGS/JunminLeeCourse/DataStructures/queues"
	"fmt"
)

func main() {
	s := queues.Queue{}
	fmt.Println(s)

	s.Enqueue(12)
	s.Enqueue(20)
	s.Enqueue(18)
	fmt.Println(s)

	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
}
