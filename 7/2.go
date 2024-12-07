package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
	"strconv"
)

func main() {
	numbers := helpers.ExtractAllNumbers(helpers.ReadAllLinesFromFile("./7/input.txt"))

	var sum int
	for _, line := range numbers {
		result := line[0]
		if _, ok := combine(line[1], line[2:], result); ok {
			sum += result
			fmt.Println("ok", line)
		}
	}
	fmt.Println(sum)
}

func combine(total int, ints []int, result int) (int, bool) {
	if len(ints) == 1 {
		addition := total + ints[0]
		multiplication := total * ints[0]
		concatenation := concat(total, ints[0])

		if addition == result || multiplication == result || concatenation == result {
			return result, true
		} else {
			return -1, false
		}
	}

	multiplication, ok := combine(total*ints[0], ints[1:], result)
	if ok {
		return multiplication, ok
	}
	addition, ok := combine(total+ints[0], ints[1:], result)
	if ok {
		return addition, ok
	}
	return combine(concat(total, ints[0]), ints[1:], result)
}

func concat(a int, b int) int {
	concat := strconv.Itoa(a) + strconv.Itoa(b)
	atoi, err := strconv.Atoi(concat)
	if err != nil {
		panic(err)
	}
	return atoi
}
