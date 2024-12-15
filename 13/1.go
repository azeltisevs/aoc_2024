package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p Point) add(point Point) Point {
	return Point{
		p.x + point.x,
		p.y + point.y,
	}
}

func (p Point) eq(point Point) bool {
	return p.x == point.x && p.y == point.y
}

func main() {
	lines := helpers.ReadAllLinesFromFile("./13/input.txt")
	numbers := helpers.ExtractAllNumbers(lines)
	fmt.Println("numbers", numbers)

	var sum int

	for i := 0; i < len(numbers); i += 4 {
		buttonAX := numbers[i][0]
		buttonAY := numbers[i][1]
		buttonBX := numbers[i+1][0]
		buttonBY := numbers[i+1][1]

		prizeX, prizeY := numbers[i+2][0], numbers[i+2][1]

		a := Point{buttonAX, buttonAY}
		b := Point{buttonBX, buttonBY}
		prize := Point{prizeX, prizeY}
		aMul, bMul := findMultipliers(a, b, prize)
		sum += findCheapestWay(a, aMul, b, bMul)

		//playSum, ok := play(Point{}, Point{buttonAX, buttonAY}, Point{buttonBX, buttonBY}, Point{prizeX, prizeY}, 0)
		//if ok {
		//	sum += playSum
		//}
	}

	fmt.Println(sum)
}

func findCheapestWay(a Point, aMul int, b Point, bMul int) int {
	cheapestCost := aMul*aCost + bMul*bCost
	lcm := helpers.LCM(a.x, b.x)
	timesA := lcm / a.x
	timesB := lcm / b.x

	for aMul > 0 {
		newCost := aMul*aCost + bMul*bCost
		fmt.Printf("Variant: %d*%d + %d*%d = %d (cost: %d)\n", a.x, aMul, b.x, bMul, a.x*aMul+b.x*bMul, newCost)
		if newCost < cheapestCost {
			cheapestCost = newCost
		}
		aMul -= timesA
		bMul += timesB
	}
	return cheapestCost
}

func findMultipliers(a Point, b Point, prize Point) (int, int) {
	aMultiplier := prize.x / a.x

	for (prize.x-a.x*aMultiplier)%b.x != 0 && aMultiplier > 0 {
		aMultiplier--
	}
	if aMultiplier == 0 {
		panic("handle me later")
	}

	return aMultiplier, (prize.x - aMultiplier*a.x) / b.x
}

// todo don't forget about 100 limit

const aCost = 3
const bCost = 1

//func play(current Point, a Point, b Point, prize Point, times int) (int, bool) {
//	lcmX := helpers.LCM(a.x, b.x)
//	prize.x / lcmX
//
//	//fmt.Println("Button A", "X: +", a.x, "; Y: +", a.y)
//	//fmt.Println("Button B", "X: +", b.x, "; Y: +", b.y)
//	//fmt.Println("Prize:", "X:", prize.x, "Y:", prize.y)
//	if current.eq(prize) {
//		return 0, true
//	}
//
//	if times > 100 {
//		return 0, false
//	}
//
//	pressA, okA := play(current.add(a), a, b, prize, times+1)
//	pressA += aCost
//	pressB, okB := play(current.add(b), a, b, prize, times+1)
//	pressB += bCost
//
//	if okA == okB {
//		return helpers.Min(pressA, pressB), okA
//	} else if okA {
//		return pressA, true
//	} else {
//		return pressB, true
//	}
//}
