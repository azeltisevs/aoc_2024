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
	filesCount := 1
	for i, block := range blocks {
		for j := 0; j < block; j++ {
			if isEmptyBlock {
				diskMap[idx] = -1
			} else {
				diskMap[idx] = i / 2
			}
			idx++
		}
		filesCount++
		isEmptyBlock = !isEmptyBlock
	}
	filesCount /= 2

	fmt.Println("files count", filesCount)
	fmt.Println("disk map", diskMap)
	fmt.Println("is fragmented?", isFragmented(diskMap))

	for i := filesCount - 1; i >= 0; i-- {
		attemptDefragment(diskMap, i)
	}

	//for i := 0; i < len(blocks)/2; i++ {
	//	defragmentRightmost(&diskMap)
	//}
	//
	fmt.Println("disk map", diskMap)
	//fmt.Println("is fragmented?", isFragmented(diskMap))
	//
	fmt.Println(checksum(diskMap))
}

func attemptDefragment(diskMap []int, filename int) {
	fileStart, fileEnd := findFileOrEmptyBlock(diskMap, filename)
	fileSize := fileEnd - fileStart + 1
	if fileStart == -1 {
		//fmt.Println("file not found", filename)
		//fmt.Println(diskMap)
		panic("file not found")
	}

	emptyBlockStart, emptyBlockEnd := findFileOrEmptyBlockForSize(diskMap, -1, fileSize)
	emptyBlockSize := emptyBlockEnd - emptyBlockStart + 1
	if emptyBlockStart != -1 && emptyBlockSize >= fileSize && emptyBlockStart < fileStart {
		//fmt.Println()
		//fmt.Println("fileStart", fileStart, "fileEnd", fileEnd, "fileSize", fileSize)
		//fmt.Println("emptyBlockStart", emptyBlockStart, "emptyBlockEnd", emptyBlockEnd, "emptyBlockSize", emptyBlockSize)
		//fmt.Println("defragmenting", diskMap, "filename", filename)
		for i := 0; i < fileSize; i++ {
			diskMap[i+emptyBlockStart] = diskMap[i+fileStart]
			diskMap[i+fileStart] = -1
		}
	}
}

func findFileOrEmptyBlockForSize(diskMap []int, filename int, size int) (int, int) {
	initialIndex := 0
	for len(diskMap) > 0 {
		s, e := findFileOrEmptyBlock(diskMap, filename)
		if e-s+1 >= size {
			return initialIndex + s, initialIndex + e
		} else {
			diskMap = diskMap[e+1:]
			initialIndex += e + 1
		}
	}
	return -1, -1
}

func findFileOrEmptyBlock(diskMap []int, filename int) (int, int) {
	fileStart := slices.Index(diskMap, filename)
	if fileStart == -1 {
		return -1, -1
	}

	fileEnd := fileStart
	for i := fileStart + 1; i < len(diskMap); i++ {
		if diskMap[i] == filename {
			fileEnd++
		} else {
			break
		}
	}
	return fileStart, fileEnd
}

func checksum(diskMap []int) (result int) {
	for i, block := range diskMap {
		if block == -1 {
			continue
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
