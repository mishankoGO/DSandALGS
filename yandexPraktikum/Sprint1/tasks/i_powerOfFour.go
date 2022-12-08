package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println("False")
		return
	}
	for n != 0 {
		rest := n % 4
		if rest != 0 {
			fmt.Println("False")
			return
		}
		n -= 4
	}
	fmt.Println("True")
}
