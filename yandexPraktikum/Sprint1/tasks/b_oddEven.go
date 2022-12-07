package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func oddEven(a, b, c int) string {
	odd, even := 0, 0
	for _, elem := range []int{a, b, c} {
		if elem%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd == 3 || even == 3 {
		return "WIN"
	}
	return "FAIL"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	values := strings.Split(line, " ")
	a, _ := strconv.Atoi(values[0])
	b, _ := strconv.Atoi(values[1])
	c, _ := strconv.Atoi(values[2])

	fmt.Println(oddEven(a, b, c))
}
