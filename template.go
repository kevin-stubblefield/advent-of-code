package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func part1(input struct{}) {
}

func part2(input struct{}) {
}

func parseInput(s *bufio.Scanner) struct{} {
	for s.Scan() {
		fmt.Println(s.Text())
	}

	return struct{}{}
}

func main() {
	f, _ := os.Open("test-input.txt")

	input := parseInput(bufio.NewScanner(f))

	start := time.Now()
	part1(input)
	fmt.Printf("Part 1: %s\n", time.Since(start))

	start = time.Now()
	part2(input)
	fmt.Printf("Part 2: %s\n", time.Since(start))
}
