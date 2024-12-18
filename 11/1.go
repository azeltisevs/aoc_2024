package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

const times = 75

type Key struct {
	stone int
	depth int
}

var cache = make(map[Key]int)

func main() {
	lines := helpers.ReadAllLinesFromFile("./11/input.txt")
	stones := helpers.ExtractNumbers(lines[0])
	fmt.Println("stones", stones)

	sum := len(stones)
	for _, stone := range stones {
		sum += blink(stone, times)
	}
	fmt.Println(sum)
}

// blink calculates how many stones are created for times blinked
func blink(stone int, times int) int {
	if times == 0 {
		return 0
	}

	if result, ok := cache[Key{stone, times}]; ok {
		return result
	}

	s1, s2 := transformStone(stone)
	var stonesProduced int
	if s2 != -1 {
		stonesProduced = 1 +
			blink(s2, times-1) +
			blink(s1, times-1)
	} else {
		stonesProduced = blink(s1, times-1)
	}

	cache[Key{stone, times}] = stonesProduced

	return stonesProduced
}

func transformStone(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	} else if digits := numberOfDigits(stone); digits%2 == 0 {
		return splitNumber(stone, digits)
	} else {
		return stone * 2024, -1
	}
}

func splitNumber(number int, digits int) (int, int) {
	secondNumber := 0

	divider := pow(10, digits/2)
	firstNumber := number / (divider / 10)

	divider /= 100
	for divider > 0 {
		secondNumber += (number / divider) % 10
		secondNumber *= 10
		divider /= 10
	}

	return firstNumber, secondNumber / 10
}

func pow(x, p int) int {
	result := x
	for i := 0; i < p; i++ {
		result *= x
	}
	return result
}

func numberOfDigits(x int) int {
	result := 0
	for x > 0 {
		result++
		x /= 10
	}
	return result
}
