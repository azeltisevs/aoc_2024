package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

func main() {
	lines := helpers.ReadAllLinesFromFile("./12/input.txt")
	fmt.Println("lines", lines)
	field := helpers.FindAllAlphanumeric(lines)
	fmt.Println("field", field)
	n := len(lines)
	m := len(lines[0])
	helpers.PrintField(field, n, m)

	var sum int
	//...
	fmt.Println(sum)
}
