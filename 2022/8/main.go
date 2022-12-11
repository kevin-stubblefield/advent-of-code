package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(trees [][]int) {
	fmt.Println(calculateVisibleTrees(trees))
}

func part2(trees [][]int) {
	fmt.Println(calculateMaxScenicScore(trees))
}

func calculateVisibleTrees(trees [][]int) int {
	width := len(trees[0])
	height := len(trees)

	onEdge := width*height - ((width - 2) * (height - 2))

	var interior int
	for i, r := range trees {
		if i == 0 || i == len(trees)-1 {
			continue
		}
		for j, c := range r {
			var up, down, left, right int
			if j == 0 || j == len(r)-1 {
				continue
			}

			// go up
			for k := i - 1; k >= 0; k-- {
				if trees[k][j] >= c {
					up = 0
					break
				}

				up = 1
			}

			// go down
			for k := i + 1; k < len(trees); k++ {
				if trees[k][j] >= c {
					down = 0
					break
				}

				down = 1
			}

			// go left
			for k := j - 1; k >= 0; k-- {
				if trees[i][k] >= c {
					left = 0
					break
				}

				left = 1
			}

			// go right
			for k := j + 1; k < len(trees[i]); k++ {
				if trees[i][k] >= c {
					right = 0
					break
				}

				right = 1
			}

			if up+down+left+right > 0 {
				interior += 1
			}
		}
	}

	return onEdge + interior
}

func calculateMaxScenicScore(trees [][]int) int {
	var scenicScores []int
	for i, r := range trees {
		if i == 0 || i == len(trees)-1 {
			continue
		}
		for j, c := range r {
			if j == 0 || j == len(r)-1 {
				continue
			}
			var up, down, left, right int

			// go up
			for k := i - 1; k >= 0; k-- {
				up += 1

				if trees[k][j] >= c {
					break
				}
			}

			// go down
			for k := i + 1; k < len(trees); k++ {
				down += 1

				if trees[k][j] >= c {
					break
				}
			}

			// go left
			for k := j - 1; k >= 0; k-- {
				left += 1

				if trees[i][k] >= c {
					break
				}
			}

			// go right
			for k := j + 1; k < len(trees[i]); k++ {
				right += 1

				if trees[i][k] >= c {
					break
				}
			}

			scenicScores = append(scenicScores, up*down*left*right)
		}
	}

	var max int
	for _, score := range scenicScores {
		if score > max {
			max = score
		}
	}
	return max
}

func getTreeHeightMap(scanner *bufio.Scanner) [][]int {
	var heightMap [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, h := range strings.Split(line, "") {
			height, _ := strconv.Atoi(h)
			row = append(row, height)
		}
		heightMap = append(heightMap, row)
	}

	return heightMap
}

func printHeightMap(m [][]int) {
	for _, i := range m {
		for _, j := range i {
			fmt.Printf("%d", j)
		}
		fmt.Print("\n")
	}
}

func main() {
	f, _ := os.Open("input.txt")

	trees := getTreeHeightMap(bufio.NewScanner(f))

	start := time.Now()
	part1(trees)
	fmt.Printf("Part 1: %s\n", time.Since(start))

	start = time.Now()
	part2(trees)
	fmt.Printf("Part 2: %s\n", time.Since(start))
}
