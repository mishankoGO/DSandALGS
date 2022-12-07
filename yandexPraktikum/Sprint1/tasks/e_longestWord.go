package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var L, ans int
	fmt.Scan(&L)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	sentense := strings.Split(line, " ")

	for _, word := range sentense {
		if len(word) > ans {
			ans = len(word)
		}
	}
	fmt.Println(ans)
}
