package main

import "fmt"

// O(n)

func findMax(arr [10]int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func findMin(arr [10]int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(findMax(arr))
	fmt.Println(findMin(arr))
}
