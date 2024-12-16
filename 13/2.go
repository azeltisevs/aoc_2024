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
		aXMul, bXMul := findMultipliers(a.x, b.x, prize.x)
		if aXMul == -1 && bXMul == -1 {
			continue
		}
		xWays := findAllWays(a.x, aXMul, b.x, bXMul)

		aYMul, bYMul := findMultipliers(a.y, b.y, prize.y)
		if aYMul == -1 && bYMul == -1 {
			continue
		}
		yWays := findAllWays(a.y, aYMul, b.y, bYMul)

		intersection := intersect(xWays, yWays)
		if len(intersection) != 0 {
			cheapest := 99999999999999
			for _, v := range intersection {
				if cheapest > v {
					cheapest = v
				}
			}
			sum += cheapest
		}

		//playSum, ok := play(Point{}, Point{buttonAX, buttonAY}, Point{buttonBX, buttonBY}, Point{prizeX, prizeY}, 0)
		//if ok {
		//	sum += playSum
		//}
	}

	fmt.Println(sum)
}

func intersect(s1 map[Mul]int, s2 map[Mul]int) map[Mul]int {
	result := make(map[Mul]int)
	for k, v := range s1 {
		if _, ok := s2[k]; ok {
			result[k] = v
		}
	}
	return result
}

type Mul struct {
	a int
	b int
}

func findAllWays(a int, aMul int, b int, bMul int) map[Mul]int {
	lcm := helpers.LCM(a, b)
	timesA := lcm / a
	timesB := lcm / b

	result := make(map[Mul]int)
	currentCost := aMul*aCost + bMul*bCost
	for aMul > 0 {
		result[Mul{aMul, bMul}] = currentCost
		fmt.Printf("Variant: %d*%d + %d*%d = %d (cost: %d)\n", a, aMul, b, bMul, a*aMul+b*bMul, currentCost)
		aMul -= timesA
		bMul += timesB
		newCost := aMul*aCost + bMul*bCost
		if newCost > currentCost {
			return result
		}
		currentCost = newCost
	}
	return result
}

func findMultipliers(a int, b int, prize int) (int, int) {
	aMultiplier := prize / a

	for (prize-a*aMultiplier)%b != 0 && aMultiplier > 0 {
		aMultiplier--
	}
	if aMultiplier == 0 {
		return -1, -1
		//panic("handle me later")
	}

	return aMultiplier, (prize - aMultiplier*a) / b
}

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
