package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackEffective struct {
	data []int
	max  int
}

func (s *StackEffective) push(x int) {
	s.data = append(s.data, x)
	if s.data[len(s.data)-1] > s.max {
		s.max = s.data[len(s.data)-1]
	}
}

func (s *StackEffective) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *StackEffective) getMax() interface{} {
	if len(s.data) == 0 {
		return "None"
	}

	return s.max
}

func OperationG(path string) {
	var stack StackEffective

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
