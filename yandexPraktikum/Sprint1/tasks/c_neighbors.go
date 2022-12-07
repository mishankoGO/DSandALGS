package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func neighbors(matrix [][]int, tRow, tCol int) string {
	var neighbor []string
	n, m := len(matrix)-1, len(matrix[0])-1
	if tRow != 0 && tRow != n {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow-1][tCol]))
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow+1][tCol]))
	} else if tRow == n {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow-1][tCol]))
	} else if tRow == 0 {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow+1][tCol]))
	}

	if tCol != 0 && tCol != m {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow][tCol-1]))
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow][tCol+1]))
	} else if tCol == m {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow][tCol-1]))
	} else if tCol == 0 {
		neighbor = append(neighbor, strconv.Itoa(matrix[tRow][tCol+1]))
	}
	neighborStr := strings.Join(neighbor, " ")
	return neighborStr
}

func main() {
	var ans string
	var n, m, tRow, tCol int
	fmt.Scan(&n)
	fmt.Scan(&m)

	var matrix = make([][]int, n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		row := func(values []string) []int {
			var row []int
			for _, val := range values {
				valInt, _ := strconv.Atoi(val)
				row = append(row, valInt)
			}
			return row
		}(values)
		matrix[i] = row
	}
	fmt.Println(matrix)
	fmt.Scan(&tRow)
	fmt.Scan(&tCol)

	ans = neighbors(matrix, tRow, tCol)
	fmt.Println(ans)
}
