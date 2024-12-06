package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const GuardUp = '^'
const GuardDown = 'V'
const GuardLeft = '<'
const GuardRight = '>'
const Crate = '#'
const EmptySpace = '.'

func main() {

	lines := helpers.ReadAllLinesFromFile("./6/input.txt")

	var field [][]rune

	path := make([][]int, len(lines), len(lines))
	for i, line := range lines {
		field = append(field, []rune(line))
		path[i] = make([]int, len(field[i]), len(field[i]))
	}

	fmt.Println("field", "\n"+drawField(field))

	y, x := locateGuard(field)
	fmt.Println("initial path", "\n"+drawField(path))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered panic", r)

			fmt.Println("x", x, "y", y)
			fmt.Println("panic field", "\n"+drawField(field))
			fmt.Println("panic path", "\n"+drawField(path))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Println("caught ^C", sig)

			fmt.Println("x", x, "y", y)
			fmt.Println("panic field", "\n"+drawField(field))
			fmt.Println("panic path", "\n"+drawField(path))

			os.Exit(130)
		}
	}()

	for ; x != -1 && y != -1; y, x = locateGuard(field) {
		path[y][x]++
		if obstacle, end := hasObstacle(field, y, x); obstacle {
			rotate(field, y, x)
		} else if end {
			break
		} else {
			move(field, y, x)
		}
	}

	fmt.Println("field", "\n"+drawField(field))
	fmt.Println("final path", "\n"+drawField(path))

	positions := countPositions(path)
	fmt.Println("unique positions", positions)

}

func countPositions(path [][]int) int {
	result := 0
	for y := range path {
		for x := range path[y] {
			if path[y][x] > 0 {
				result++
			}
		}
	}
	return result
}

func move(field [][]rune, y int, x int) {
	switch field[y][x] {
	case GuardUp:
		field[y-1][x] = GuardUp
	case GuardDown:
		field[y+1][x] = GuardDown
	case GuardLeft:
		field[y][x-1] = GuardLeft
	case GuardRight:
		field[y][x+1] = GuardRight
	}
	field[y][x] = EmptySpace
}

func rotate(field [][]rune, y int, x int) {
	switch field[y][x] {
	case GuardUp:
		field[y][x] = GuardRight
	case GuardDown:
		field[y][x] = GuardLeft
	case GuardLeft:
		field[y][x] = GuardUp
	case GuardRight:
		field[y][x] = GuardDown
	}
}

func hasObstacle(field [][]rune, y int, x int) (obstacle bool, end bool) {
	switch field[y][x] {
	case GuardUp:
		if y-1 < 0 {
			return false, true
		}
		return field[y-1][x] == Crate, false
	case GuardDown:
		if y+1 >= len(field[y]) {
			return false, true
		}
		return field[y+1][x] == Crate, false
	case GuardLeft:
		if x-1 < 0 {
			return false, true
		}
		return field[y][x-1] == Crate, false
	case GuardRight:
		if x+1 >= len(field) {
			return false, true
		}
		return field[y][x+1] == Crate, false
	}

	panic("guard not identified " + fmt.Sprintf("[%d %d]: %v", y, x, field[y][x]))
}

func drawField[V any](field [][]V) string {
	result := ""
	for _, row := range field {
		for _, v := range row {
			if _, ok := any(v).(rune); ok {
				result += fmt.Sprintf("%c ", v)
			} else {
				result += fmt.Sprintf("%v ", v)
			}
		}
		result += fmt.Sprintf("\n")
	}
	return result
}

func locateGuard(field [][]rune) (int, int) {
	for y := range field {
		for x := range field[y] {
			if field[y][x] == GuardRight ||
				field[y][x] == GuardLeft ||
				field[y][x] == GuardDown ||
				field[y][x] == GuardUp {
				return y, x
			}
		}
	}
	return -1, -1
}
