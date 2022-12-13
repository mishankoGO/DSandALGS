package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackMax struct {
	data []int
}

func (s *StackMax) push(x int) {
	s.data = append(s.data, x)
}

func (s *StackMax) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *StackMax) getMax() interface{} {
	if len(s.data) == 0 {
		return "None"
	}
	var max = s.data[0]

	for _, elem := range s.data {
		if elem > max {
			max = elem
		}
	}
	return max
}

func Operation(path string) {
	var stack StackMax

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimRight(scanner.Text(), "\n"))

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		command := strings.Split(line, " ")
		switch command[0] {
		case "getMax":
			fmt.Println(stack.getMax())
		case "pop":
			stack.pop()
		case "push":
			val, _ := strconv.Atoi(command[1])
			stack.push(val)
		}
	}
}
func main() {
	Operation("f_test.txt")
}
