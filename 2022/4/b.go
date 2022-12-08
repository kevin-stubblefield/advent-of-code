package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}

	answer := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		answer += process(scanner.Text())
	}

	fmt.Printf("The answer is %d", answer)
}

func process(line string) int {
	assignments := strings.Split(line, ",")
	pair1 := strings.Split(assignments[0], "-")
	pair2 := strings.Split(assignments[1], "-")
	a, _ := strconv.Atoi(pair1[0])
	b, _ := strconv.Atoi(pair1[1])
	c, _ := strconv.Atoi(pair2[0])
	d, _ := strconv.Atoi(pair2[1])

	if a <= c && b >= d {
		return 1
	}

	if a >= c && b <= d {
		return 1
	}

	if a >= c && a <= d {
		return 1
	}

	if b >= c && b <= d {
		return 1
	}

	return 0
}
