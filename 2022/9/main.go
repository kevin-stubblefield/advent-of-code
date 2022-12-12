package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type knot struct {
	x int
	y int
}

func areTouching(hX, hY, tX, tY int) bool {
	if math.Abs(float64(hX-tX)) <= 1 && int(math.Abs(float64(hY-tY))) <= 1 {
		return true
	}

	return false
}

func moveHeadOneStep(x, y int, direction string) (int, int) {
	switch direction {
	case "U":
		return x, y + 1
	case "D":
		return x, y - 1
	case "L":
		return x - 1, y
	case "R":
		return x + 1, y
	default:
		return x, y
	}
}

func moveTail(hX, hY, tX, tY int) (int, int) {
	var dX, dY int
	dX = hX - tX
	dY = hY - tY

	absDX := math.Abs(float64(dX))
	absDY := math.Abs(float64(dY))

	if absDX == 2 && dY == 0 {
		if dX < 0 {
			dX = -1
		} else {
			dX = 1
		}
	} else if absDY == 2 && dX == 0 {
		if dY < 0 {
			dY = -1
		} else {
			dY = 1
		}
	} else if absDX+absDY > 2 && absDX*absDY != 0 {
		if dX < 0 {
			dX = -1
		} else {
			dX = 1
		}

		if dY < 0 {
			dY = -1
		} else {
			dY = 1
		}
	}

	return tX + dX, tY + dY
}

func part1(s *bufio.Scanner) {
	var hX, hY, tX, tY int

	tailVisited := make(map[string]bool)
	tailVisited[fmt.Sprintf("%d,%d", tX, tY)] = true

	for s.Scan() {
		tokens := strings.Split(s.Text(), " ")
		steps, _ := strconv.Atoi(tokens[1])
		for i := 0; i < steps; i++ {
			// fmt.Printf("HEAD: (%d, %d)\n", hX, hY)
			// fmt.Printf("TAIL: (%d, %d)\n", tX, tY)
			hX, hY = moveHeadOneStep(hX, hY, tokens[0])
			if !areTouching(hX, hY, tX, tY) {
				tX, tY = moveTail(hX, hY, tX, tY)
				tailVisited[fmt.Sprintf("%d,%d", tX, tY)] = true
			}
		}
	}

	fmt.Printf("Answer: %d\n", len(tailVisited))
}

func part2(s *bufio.Scanner) {
	tailVisited := make(map[string]bool)
	tailVisited[fmt.Sprintf("%d,%d", 0, 0)] = true

	var knots []knot
	for i := 0; i < 10; i++ {
		knots = append(knots, knot{x: 0, y: 0})
	}

	for s.Scan() {
		tokens := strings.Split(s.Text(), " ")
		steps, _ := strconv.Atoi(tokens[1])
		for i := 0; i < steps; i++ {
			// fmt.Printf("HEAD: (%d, %d)\n", hX, hY)
			// fmt.Printf("TAIL: (%d, %d)\n", tX, tY)
			knots[0].x, knots[0].y = moveHeadOneStep(knots[0].x, knots[0].y, tokens[0])
			for k := 1; k < 10; k++ {
				if !areTouching(knots[k-1].x, knots[k-1].y, knots[k].x, knots[k].y) {
					knots[k].x, knots[k].y = moveTail(knots[k-1].x, knots[k-1].y, knots[k].x, knots[k].y)
				}
			}
			tailVisited[fmt.Sprintf("%d,%d", knots[9].x, knots[9].y)] = true
		}
	}

	fmt.Printf("Answer: %d\n", len(tailVisited))
}

func parseInput(s *bufio.Scanner) struct{} {
	for s.Scan() {
		fmt.Println(s.Text())
	}

	return struct{}{}
}

func main() {
	f, _ := os.Open("input.txt")

	start := time.Now()
	part1(bufio.NewScanner(f))
	fmt.Printf("Part 1: %s\n", time.Since(start))

	f, _ = os.Open("input.txt")
	start = time.Now()
	part2(bufio.NewScanner(f))
	fmt.Printf("Part 2: %s\n", time.Since(start))
}
