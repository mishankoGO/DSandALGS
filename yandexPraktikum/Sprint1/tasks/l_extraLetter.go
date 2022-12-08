package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	sVals := strings.Split(s, "")
	tVals := strings.Split(t, "")

	var count = make(map[string]int)

	for _, elem := range sVals {
		count[elem] += 1
	}

	for _, elem := range tVals {
		if _, ok := count[elem]; ok {
			count[elem] = count[elem] - 1
			if count[elem] == 0 {
				delete(count, elem)
			}
		} else {
			fmt.Println(elem)
		}
	}
}
