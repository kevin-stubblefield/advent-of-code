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

func part1(s *bufio.Scanner) {
	fmt.Printf("Answer: %d\n", solvePart1(s, 20, 40))
}

func part2(s *bufio.Scanner) {
	solvePart2(s)
}

func solvePart1(s *bufio.Scanner, significantCycleStart, significantCycleIncrement int) int {
	answer := 0
	cycle := 0
	v := 0
	register := 1
	nextSignificantCycle := significantCycleStart

	for s.Scan() {
		numCycles := 1
		tokens := strings.Split(s.Text(), " ")
		if tokens[0] == "addx" {
			numCycles = 2
			v, _ = strconv.Atoi(tokens[1])
		}

		for i := 0; i < numCycles; i++ {
			cycle += 1
			if cycle == nextSignificantCycle {
				answer += cycle * register
				nextSignificantCycle += significantCycleIncrement
			}
		}

		if cycle > 220 {
			break
		}

		if numCycles == 2 {
			register += v
		}
	}

	return answer
}

func solvePart2(s *bufio.Scanner) {
	var crt [6][40]string

	cycle := 0
	v := 0
	register := 1
	row := 0

	for s.Scan() {
		numCycles := 1
		tokens := strings.Split(s.Text(), " ")
		if tokens[0] == "addx" {
			numCycles = 2
			v, _ = strconv.Atoi(tokens[1])
		}

		for i := 0; i < numCycles; i++ {
			drawPos := cycle % 40
			cycle += 1

			// fmt.Printf("Cycle: %d, DrawPos: %d, Register: %d\n", cycle, drawPos, register)

			if drawPos >= register-1 && drawPos <= register+1 {
				crt[row][drawPos] = "#"
			} else {
				crt[row][drawPos] = "."
			}

			row = int(math.Floor(float64(cycle / 40)))
		}

		if cycle > 240 {
			break
		}

		if numCycles == 2 {
			register += v
		}
	}

	for i := 0; i < cap(crt); i++ {
		for j := 0; j < cap(crt[i]); j++ {
			fmt.Print(crt[i][j])
		}
		fmt.Print("\n")
	}
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
