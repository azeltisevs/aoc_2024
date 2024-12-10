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
		trails := findTrails(field, point, 0)
		fmt.Println("trails", trails)
		sum += len(trails)
	}
	fmt.Println("sum", sum)
}

func findTrails(field [][]int, point Point, next int) map[Point]bool {
	if point.y < 0 || point.x < 0 || point.y >= len(field) || point.x >= len(field[point.y]) {
		return nil
	}

	if field[point.y][point.x] != next {
		return nil
	} else if next == 9 {
		result := make(map[Point]bool)
		result[point] = true
		return result
	}

	result := make(map[Point]bool)
	mergeMap(result, findTrails(field, Point{point.y - 1, point.x}, next+1))
	mergeMap(result, findTrails(field, Point{point.y, point.x + 1}, next+1))
	mergeMap(result, findTrails(field, Point{point.y + 1, point.x}, next+1))
	mergeMap(result, findTrails(field, Point{point.y, point.x - 1}, next+1))
	return result
}

func mergeMap(result map[Point]bool, trails map[Point]bool) {
	for k, v := range trails {
		result[k] = v
	}
}
