package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := helpers.ReadAllLinesFromFile("./5/input.txt")

	var order = make(map[int][]int)

	i := 0
	for lines[i] != "" {
		line := lines[i]
		split := strings.Split(line, "|")

		first, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		then, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		if _, ok := order[first]; !ok {
			// Init map entry
			order[first] = []int{then}
		} else {
			order[first] = append(order[first], then)
		}

		i++
	}

	fmt.Println("order", order)
	result := 0
	for i++; i < len(lines); i++ {
		split := strings.Split(lines[i], ",")
		var updates []int
		for _, updateString := range split {
			update, err := strconv.Atoi(updateString)
			if err != nil {
				panic(err)
			}
			updates = append(updates, update)
		}

		if isCorrectUpdate(updates, order) {
			fmt.Println("update", updates, "is correct, adding", updates[len(updates)/2])
			result += updates[len(updates)/2]
		}
	}

	fmt.Println(result)
}

func isCorrectUpdate(update []int, order map[int][]int) bool {
	for i := len(update) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(order[update[i]], update[j]) {
				return false
			}
		}
	}
	return true
}
