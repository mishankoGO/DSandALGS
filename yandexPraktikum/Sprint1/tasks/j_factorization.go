package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var ans []string

	// Get the number of 2s that divide n
	for n%2 == 0 {
		ans = append(ans, strconv.Itoa(2))
		n /= 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			ans = append(ans, strconv.Itoa(i))
			n /= i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		ans = append(ans, strconv.Itoa(n))
	}

	fmt.Println(strings.Join(ans, " "))
}
