package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
	"slices"
)

func main() {
	line := helpers.ReadAllLinesFromFile("./9/input.txt")[0]
	blocks := helpers.ExtractDigits(line)
	fmt.Println("blocks", blocks)

	var diskSize int
	for _, block := range blocks {
		diskSize += block
	}

	fmt.Println("disk size", diskSize)
	diskMap := make([]int, diskSize)
	//for i := range diskMap {
	//	diskMap[i] = -1
	//}
	//for i := range numbers {
	//	diskMap[i] = numbers[i]
	//}

	idx := 0
	isEmptyBlock := false
	for i, block := range blocks {
		for j := 0; j < block; j++ {
			if isEmptyBlock {
				diskMap[idx] = -1
			} else {
				diskMap[idx] = i / 2
			}
			idx++
		}
		isEmptyBlock = !isEmptyBlock
	}

	fmt.Println("disk map", diskMap)
	fmt.Println("is fragmented?", isFragmented(diskMap))

	for isFragmented(diskMap) {
		defragmentRightmost(&diskMap)
	}

	fmt.Println("disk map", diskMap)
	fmt.Println("is fragmented?", isFragmented(diskMap))

	fmt.Println(checksum(diskMap))
}

func checksum(diskMap []int) (result int) {
	for i, block := range diskMap {
		if block == -1 {
			return
		}
		result += i * block
	}
	return
}

func defragmentRightmost(diskMap *[]int) {
	for i := len(*diskMap) - 1; i >= 1; i-- {
		if (*diskMap)[i] != -1 {
			emptyBlockIdx := slices.Index(*diskMap, -1)
			(*diskMap)[i], (*diskMap)[emptyBlockIdx] = (*diskMap)[emptyBlockIdx], (*diskMap)[i]
			return
		}
	}
}

func isFragmented(diskMap []int) bool {
	firstEmptyBlockIndex := slices.Index(diskMap, -1)

	for i := firstEmptyBlockIndex + 1; i < len(diskMap); i++ {
		if diskMap[i] != -1 {
			return true
		}
	}
	return false
}
