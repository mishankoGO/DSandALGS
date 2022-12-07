package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveEquation(a, b, c, x int) int {
	ans := a*x*x + b*x + c
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	values := strings.Split(line, " ")
	a, _ := strconv.Atoi(values[0])
	x, _ := strconv.Atoi(values[1])
	b, _ := strconv.Atoi(values[2])
	c, _ := strconv.Atoi(values[3])
	fmt.Println(solveEquation(a, b, c, x))
}
