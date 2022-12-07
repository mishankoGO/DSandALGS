// Distance of closest zero to every element
// answer: https://www.geeksforgeeks.org/distance-closest-zero-every-element/
// Time: O(2n), Space: O(3n)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper function to find min distance from 2 arrays
func findMin(arr1 []int, arr2 []int) string {
	var res []string
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == arr2[i] {
			res = append(res, strconv.Itoa(arr1[i]))
		} else if arr1[i] < arr2[i] {
			if arr1[i] == 0 { // there cannot be zero when element in other array is not zero
				res = append(res, strconv.Itoa(arr2[i]))
			} else {
				res = append(res, strconv.Itoa(arr1[i]))
			}
		} else {
			if arr2[i] == 0 {
				res = append(res, strconv.Itoa(arr1[i]))
			} else {
				res = append(res, strconv.Itoa(arr2[i]))
			}
		}
	}
	return strings.Join(res, " ")
}

// function to find the solution
func nearestNeighbor(n int, values []string) string {
	// create two arrays for forward and backward passes
	var ans1, ans2 = make([]int, n), make([]int, n)

	// if there is only one place, it should be empty
	if n == 1 {
		return "0"
	}

	// first we pass through the houses in forward manner
	// and if we stumble a 0, then we start counting houses until next 0
	// and store that counter in the ans1
	flag := 0
	k := 0
	for i := 0; i < n; i++ {
		valInt, _ := strconv.Atoi(values[i])
		if valInt == 0 {
			ans1[i] = 0
			k = 0
			flag = 1
		} else {
			ans1[i] = k
		}
		if flag == 1 {
			k++
		}
	}

	// then we run through values in backward manner
	// and when we stumble a 0 then we start counting houses until next blank space
	k = 0
	flag = 0
	for i := n - 1; i >= 0; i-- {
		valInt, _ := strconv.Atoi(values[i])
		if valInt == 0 {
			ans2[i] = 0
			k = 0
			flag = 1
		} else {
			ans2[i] = k
		}
		if flag == 1 {
			k++
		}
	}

	// find the minimum value for each position
	ans := findMin(ans1, ans2)
	return ans
}

func main() {

	//tests := make(map[string]string)
	//tests["0 1 4 9 0"] = "0 1 2 1 0"
	//tests["0 7 9 4 8 20"] = "0 1 2 3 4 5"
	//tests["0 1 2 3 0 1 2 0 1 2"] = "0 1 2 1 0 1 1 0 1 2"
	//tests["1 2 3 0 1 2 0 1 2 0"] = "3 2 1 0 1 1 0 1 1 0"
	//tests["1 2 3 0 1 2 3"] = "3 2 1 0 1 2 3"
	//tests["0 2 3 0 1 2 3"] = "0 1 1 0 1 2 3"
	//tests["0 0"] = "0 0"
	//tests["0 1 0"] = "0 1 0"
	//tests["1 1 1 1 1 0"] = "5 4 3 2 1 0"
	//tests["2 1 0 3 0 0 3 2 4"] = "2 1 0 1 0 0 1 2 3"
	//tests["5 6 0 1 -2 3 4"] = "2 1 0 1 2 3 4"
	//tests["2 1 0 3 0 0 3 2 4"] = "2 1 0 1 0 0 1 2 3"
	//
	//for k, v := range tests {
	//	values := strings.Split(k, " ")
	//	n := len(values)
	//	ans := nearestNeighbor(n, values)
	//	fmt.Println(ans)
	//	if ans != v {
	//		fmt.Printf("failed test %s\n", k)
	//		fmt.Printf("expected %s, got %s\n", v, ans)
	//		panic("tests don't pass")
	//	}
	//}
	//fmt.Println("all tests passed!")

	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	line := scanner.Text()
	values := strings.Split(line, " ")
	ans := nearestNeighbor(n, values)
	fmt.Println(ans)
}
