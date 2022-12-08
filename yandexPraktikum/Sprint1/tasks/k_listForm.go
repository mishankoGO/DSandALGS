package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, K int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	values := strings.ReplaceAll(line, " ", "")

	X, _ := strconv.Atoi(values)
	fmt.Scan(&K)
	ans := strconv.Itoa(X + K)
	fmt.Println(strings.Join(strings.Split(ans, ""), " "))
}
