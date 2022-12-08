package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processB(group []string, priority string) int {
	var shared byte

	for i := range group[0] {
		found := strings.IndexByte(group[1], group[0][i])
		if found > -1 {
			foundAgain := strings.IndexByte(group[2], group[0][i])
			if foundAgain > -1 {
				shared = group[0][i]
				break
			}
		}
	}

	return strings.IndexByte(priority, shared)
}

func main() {
	fmt.Println("Solution for part b")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Couldn't open file")
	}

	priority := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	answer := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var group []string
		group = append(group, scanner.Text())
		scanner.Scan()
		group = append(group, scanner.Text())
		scanner.Scan()
		group = append(group, scanner.Text())
		answer += processB(group, priority)
	}

	fmt.Printf("The answer is %d\n", answer)
}
