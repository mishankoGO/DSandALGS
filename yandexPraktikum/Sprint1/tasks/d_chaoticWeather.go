package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, res int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	if n == 1 {
		fmt.Println(scanner.Text())
		return
	}

	line := scanner.Text()
	values := strings.Split(line, " ")

	temps := func(values []string) (res []int) {
		for _, elem := range values {
			elemInt, _ := strconv.Atoi(elem)
			res = append(res, elemInt)
		}
		return res
	}(values)

	if temps[0] > temps[1] {
		res++
	}
	if temps[n-1] > temps[n-2] {
		res++
	}
	for i := 1; i < len(temps)-1; i++ {
		curr := temps[i]
		if curr > temps[i-1] && curr > temps[i+1] {
			res++
		}
	}
	fmt.Println(res)
}
