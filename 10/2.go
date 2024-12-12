package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

type Point struct {
	y int
	x int
}

func main() {
	field := helpers.ExtractAllDigits(helpers.ReadAllLinesFromFile("./10/input.txt"))

	var trailheads []Point
	for y := range field {
		for x := range field[y] {
			if field[y][x] == 0 {
				trailheads = append(trailheads, Point{y, x})
			}
		}
	}

	fmt.Println("trailheads", trailheads)

	var sum int
	for _, point := range trailheads {
		sum += findTrails(field, point, 0)
	}
	fmt.Println("sum", sum)
}

func findTrails(field [][]int, point Point, next int) int {
	if point.y < 0 || point.x < 0 || point.y >= len(field) || point.x >= len(field[point.y]) {
		return 0
	}

	if field[point.y][point.x] != next {
		return 0
	} else if next == 9 {
		return 1
	}

	return findTrails(field, Point{point.y - 1, point.x}, next+1) +
		findTrails(field, Point{point.y, point.x + 1}, next+1) +
		findTrails(field, Point{point.y + 1, point.x}, next+1) +
		findTrails(field, Point{point.y, point.x - 1}, next+1)
}
