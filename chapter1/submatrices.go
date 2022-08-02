package main

import "fmt"

func findSubmatrices(matrix [][]int) [7][2][2]int {
	var res [7][2][2]int

	var coords = [][][]int{
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		{{0, 0}, {0, 1}, {2, 0}, {2, 1}},
		{{1, 0}, {1, 1}, {2, 0}, {2, 1}},
		{{0, 1}, {0, 2}, {1, 1}, {1, 2}},
		{{0, 1}, {0, 2}, {2, 1}, {2, 2}},
		{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
		{{0, 0}, {0, 2}, {2, 0}, {2, 2}},
	}
	for i, coord := range coords {
		res[i][0][0] = matrix[coord[0][0]][coord[0][1]]
		res[i][0][1] = matrix[coord[1][0]][coord[1][1]]
		res[i][1][0] = matrix[coord[2][0]][coord[2][1]]
		res[i][1][1] = matrix[coord[3][0]][coord[3][1]]
	}
	return res
}

func main() {
	subMat := findSubmatrices([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	for _, elem := range subMat {
		fmt.Println(elem)
	}
}
