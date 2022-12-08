package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var res []string
	for n != 0 {
		rest := n % 2
		n /= 2
		res = append([]string{strconv.Itoa(rest)}, res...)
	}
	fmt.Println(strings.Join(res, ""))
}
