package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Предлагается создать словарь всех вхождений каждого символа.
// Затем пробежаться по словарю и если число значений меньше или равно k*2
// то записываем балл и вычитаем занятые пальцы

func agileHandsAtOnce(vals map[int]int, k int) int {
	var answers []int
	for k1, v1 := range vals { // пробегаемся по максимум 9 значениям
		fingers := k * 2
		buf := 0
		if v1 <= fingers {
			buf += 1
			fingers -= v1
		}
		for k2, v2 := range vals { // пробегаемся по максимум 9 значениям (в купе с первым циклом максимум 81 прогон)
			if v2 > fingers || k1 == k2 {
				answers = append(answers, buf)
				continue
			}
			fingers -= v2
			buf += 1
		}
	}
	sort.Ints(answers) // если сортировка быстрая, то nlogn
	return answers[len(answers)-1]
}

func agileHands(vals map[int]int, k int) int {
	var ans int
	fingers := k * 2
	for _, v := range vals {
		if v <= fingers {
			ans++
		}
	}
	return ans
}

func main() {
	var k int
	var vals = make(map[int]int) // хранит максимум 9 ключей и максимальное значение 16
	fmt.Scan(&k)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()
		for _, elem := range line {
			if string(elem) != "." {
				elemInt, _ := strconv.Atoi(string(elem))
				vals[elemInt] += 1
			}
		}
	}
	fmt.Println(vals)

	fmt.Println("one at the time:")
	fmt.Println(agileHands(vals, k))

	fmt.Println("everything at once:")
	fmt.Println(agileHandsAtOnce(vals, k))
}
