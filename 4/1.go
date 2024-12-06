package main

import (
	"fmt"
	"strings"
)

const xmas = "XMAS"

func main() {
	var input = "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	lines := strings.Split(input, "\n")

	var counter = 0
	for _, line := range lines {
		counter += countXmas(line)
	}
	for i := range lines[0] {
		var vertical []rune
		for _, line := range lines {
			vertical = append(vertical, rune(line[i]))
		}
		counter += countXmas(string(vertical))
	}

	fmt.Println(counter)
}

func countXmas(line string) int {
	return strings.Count(line, xmas) + strings.Count(Reverse(line), xmas)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
