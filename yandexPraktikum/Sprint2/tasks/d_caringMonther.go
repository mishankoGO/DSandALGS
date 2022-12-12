package tasks

func SolutionD(head *ListNode, value string) int {
	var index int
	for node := head; node != nil; node = node.next {
		if node.data == value {
			return index
		}
		index++
	}
	return -1
}
