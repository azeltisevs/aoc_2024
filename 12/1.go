package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

type Point struct {
	y int
	x int
}

type Set struct {
	internalSet map[Point]bool
}

func (s Set) add(p Point) {
	s.internalSet[p] = true
}

func (s Set) contains(p Point) bool {
	_, ok := s.internalSet[p]
	return ok
}

func (s Set) values() []Point {
	keys := make([]Point, len(s.internalSet))

	i := 0
	for k := range s.internalSet {
		keys[i] = k
		i++
	}

	return keys
}

func main() {
	lines := helpers.ReadAllLinesFromFile("./12/input.txt")
	field := helpers.ReadAllRunes(lines)
	helpers.PrintField(field)

	// perimeter candidate is an entry in a field that has <4 neighbours matching itself

	calculated := make(map[Point]bool)
	var sum int
	for y := range field {
		for x := range field[y] {
			point := Point{y, x}
			if calculated[point] {
				continue
			} else {
				neighbourhood := Set{internalSet: make(map[Point]bool)}
				findNeighbours(field, field[y][x], point, neighbourhood)
				neighbors := neighbourhood.values()
				fmt.Println(neighbors)
				for _, p := range neighbors {
					calculated[p] = true
				}
				area := len(neighbors)
				perimeter := calculatePerimeter(field, neighbors)
				fmt.Printf("%c; Area: %d, Perimeter: %d\n", field[y][x], area, perimeter)
				sum += area * perimeter
			}
		}
	}
	fmt.Println(sum)
}

func calculatePerimeter(field [][]rune, neighbors []Point) int {
	perimeter := 0
	for _, point := range neighbors {
		neighborsCount := countNeighbors(field, point)
		perimeter += 4 - neighborsCount
	}
	return perimeter
}

func countNeighbors(field [][]rune, point Point) (result int) {
	char := field[point.y][point.x]
	if eq(field, char, Point{point.y - 1, point.x}) {
		result++
	}
	if eq(field, char, Point{point.y + 1, point.x}) {
		result++
	}
	if eq(field, char, Point{point.y, point.x - 1}) {
		result++
	}
	if eq(field, char, Point{point.y, point.x + 1}) {
		result++
	}
	return
}

func eq(field [][]rune, char rune, point Point) bool {
	y := point.y
	x := point.x
	if y < 0 ||
		x < 0 ||
		y >= len(field) ||
		x >= len(field[y]) {
		return false
	}
	return field[y][x] == char
}

func findNeighbours(field [][]rune, char rune, point Point, neighbourhood Set) {
	y := point.y
	x := point.x
	if y < 0 ||
		x < 0 ||
		y >= len(field) ||
		x >= len(field[y]) {
		return
	}
	if field[y][x] != char {
		return
	}

	neighbourhood.add(point)

	if direction := (Point{y + 1, x}); !neighbourhood.contains(direction) {
		findNeighbours(field, char, direction, neighbourhood)
	}
	if direction := (Point{y - 1, x}); !neighbourhood.contains(direction) {
		findNeighbours(field, char, direction, neighbourhood)
	}
	if direction := (Point{y, x + 1}); !neighbourhood.contains(direction) {
		findNeighbours(field, char, direction, neighbourhood)
	}
	if direction := (Point{y, x - 1}); !neighbourhood.contains(direction) {
		findNeighbours(field, char, direction, neighbourhood)
	}
}
