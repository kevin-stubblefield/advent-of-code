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
	first, _ := strconv.Atoi(pair1[0])
	second, _ := strconv.Atoi(pair1[1])
	third, _ := strconv.Atoi(pair2[0])
	fourth, _ := strconv.Atoi(pair2[1])

	if first <= third && second >= fourth {
		return 1
	}

	if first >= third && second <= fourth {
		return 1
	}

	return 0
}
