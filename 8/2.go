package main

import (
	"advent_of_code_2024/helpers"
	"fmt"
)

func main() {
	lines := helpers.ReadAllLinesFromFile("./8/input.txt")
	n, m := len(lines), len(lines[0])
	antennas := helpers.FindAllAlphanumericGroupByRunes(lines)

	helpers.PrintFieldMap(antennas, n, m)

	antinodes := make(map[helpers.Point]int)
	for antenna, points := range antennas {
		fmt.Println("antenna", antenna)
		for i := range points {
			antinodes[points[i]]++
			for j := i + 1; j < len(points); j++ {
				dx1, dy1 := points[i].X-points[j].X, points[i].Y-points[j].Y
				dx2, dy2 := points[j].X-points[i].X, points[j].Y-points[i].Y

				fmt.Println(points[i], points[j])
				fmt.Println("dx1", dx1, "dy1", dy1)
				fmt.Println("dx2", dx2, "dy2", dy2)

				antinode1 := helpers.Point{points[i].Y + dy1, points[i].X + dx1}
				for checkAntinode(antinode1, m, n) {
					point := helpers.Point{
						antinode1.Y, antinode1.X,
					}
					antinodes[point]++

					antinode1.Y += dy1
					antinode1.X += dx1
				}
				antinode2 := helpers.Point{points[j].Y + dy2, points[j].X + dx2}
				for checkAntinode(antinode2, m, n) {
					point := helpers.Point{antinode2.Y, antinode2.X}
					antinodes[point]++
					antinode2.Y += dy2
					antinode2.X += dx2
				}
			}
		}

		fmt.Println(len(antinodes))
	}
}

func checkAntinode(antinode1 helpers.Point, m int, n int) bool {
	return antinode1.X >= 0 && antinode1.X < m &&
		antinode1.Y >= 0 && antinode1.Y < n
}
