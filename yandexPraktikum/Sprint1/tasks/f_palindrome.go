package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	var str, revStr string
	for _, char := range line {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			str += strings.ToLower(string(char))
		}
	}
	fmt.Println(str)

	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	if revStr == str {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
