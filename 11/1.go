package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

type Node struct {
	stone int
	next  *Node
}

func main() {
	lines := helpers.ReadAllLinesFromFile("./11/input.txt")
	stones := helpers.ExtractNumbers(lines[0])
	fmt.Println("stones", stones)

	cache := make(map[int][]int)
	cache[0] = []int{1}

	head := Node{
		stones[0],
		nil,
	}
	currentNode := &head
	for i := 1; i < len(stones); i++ {
		newNode := &Node{
			stones[i],
			nil,
		}
		currentNode.next = newNode
		currentNode = newNode
	}
	printList(&head)

	for i := 0; i < 75; i++ {
		blink(&head, cache)
		fmt.Println(i + 1)
		//printList(&head)
	}

	sum := 0
	currentNode = &head
	for currentNode != nil {
		currentNode = currentNode.next
		sum++
	}
	fmt.Println(sum)
}

func printList(head *Node) {
	currentNode := head
	fmt.Print("[ ")
	for currentNode != nil {
		fmt.Print(currentNode.stone, " ")
		currentNode = currentNode.next
	}
	fmt.Println("]")
}

func blink(head *Node, cache map[int][]int) {
	currentNode := head
	for currentNode != nil {
		if newStones, ok := cache[currentNode.stone]; ok {
			if len(newStones) > 1 {
				currentNode.stone = newStones[0]
				newNode := &Node{
					stone: newStones[1],
					next:  currentNode.next,
				}
				currentNode.next = newNode
				currentNode = newNode
			} else {
				currentNode.stone = newStones[0]
			}
		} else {
			if digits := numberOfDigits(currentNode.stone); digits%2 == 0 {
				firstPart, secondPart := splitNumber(currentNode.stone, digits)
				cache[currentNode.stone] = []int{firstPart, secondPart}

				currentNode.stone = firstPart
				newNode := &Node{
					stone: secondPart,
					next:  currentNode.next,
				}
				currentNode.next = newNode
				currentNode = newNode

			} else {
				cache[currentNode.stone] = []int{currentNode.stone * 2024}
				currentNode.stone = cache[currentNode.stone][0]
			}
		}
		currentNode = currentNode.next
	}
}

func splitNumber(number int, digits int) (int, int) {
	secondNumber := 0

	divider := pow(10, digits/2)
	firstNumber := number / (divider / 10)

	divider /= 100
	for divider > 0 {
		secondNumber += (number / divider) % 10
		secondNumber *= 10
		divider /= 10
	}

	return firstNumber, secondNumber / 10
}

func pow(x, p int) int {
	result := x
	for i := 0; i < p; i++ {
		result *= x
	}
	return result
}

func numberOfDigits(x int) int {
	result := 0
	for x > 0 {
		result++
		x /= 10
	}
	return result
}
