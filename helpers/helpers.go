package helpers

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadAllLinesFromFile(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n")
}

func ExtractAllNumbers(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		result = append(result, ExtractNumbers(line))
	}
	return result
}

func ExtractNumbers(s string) []int {
	regex, err := regexp.Compile("(\\d+)")
	if err != nil {
		panic(err)
	}

	numberStrings := regex.FindAllString(s, -1)
	var result []int
	for _, numberString := range numberStrings {
		x, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		result = append(result, x)
	}

	return result
}
