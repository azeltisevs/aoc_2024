package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ReadAllLinesFromFile(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n")
}

func ExtractAllDigits(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		result = append(result, ExtractDigits(line))
	}
	return result
}

func ExtractAllNumbers(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		result = append(result, ExtractNumbers(line))
	}
	return result
}

type Point struct {
	Y int
	X int
}

func FindAllAlphanumeric(lines []string) map[rune][]Point {
	result := make(map[rune][]Point)
	for y, line := range lines {
		for x, r := range line {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				result[r] = append(result[r], Point{y, x})
			}
		}
	}
	return result
}

func ExtractDigits(s string) []int {
	return extractNumRegex(s, "(\\d)")
}

func ExtractNumbers(s string) []int {
	return extractNumRegex(s, "(\\d+)")
}

func extractNumRegex(s string, pattern string) []int {
	regex, err := regexp.Compile(pattern)
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

func PrintField(field map[rune][]Point, n int, m int) {
	inversedField := inverseField(field)
	fmt.Println()
	fmt.Printf("%3c", ' ')
	for x := 0; x < n; x++ {
		fmt.Printf("%3d", x)
	}
	fmt.Println()
	for y := 0; y < n; y++ {
		fmt.Printf("%3d", y)
		for x := 0; x < m; x++ {
			if char, ok := inversedField[Point{y, x}]; ok {
				fmt.Printf("%3c", char)
			} else {
				fmt.Printf("%3c", '.')
			}
		}
		fmt.Println()
	}
}

func inverseField(field map[rune][]Point) map[Point]rune {
	result := make(map[Point]rune)
	for k, v := range field {
		for _, point := range v {
			//if value, ok := result[point]; ok {
			//	panic(fmt.Sprintf("conflict. key %v, existing value: %c, new value: %c", point, value, k))
			//}
			result[point] = k
		}
	}
	return result
}
