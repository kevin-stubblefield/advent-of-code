package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processA(line string, priority string) int {
	bytes := []byte(line)
	half := len(bytes) / 2
	firstHalf := bytes[:half]
	secondHalf := bytes[half:]

	var shared byte
	for i := range firstHalf {
		found := strings.IndexByte(string(secondHalf), firstHalf[i])
		if found > -1 {
			shared = firstHalf[i]
		}
	}

	return strings.IndexByte(priority, shared)
}

func main() {
	fmt.Println("Solution for part a")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Couldn't open file")
	}

	priority := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	answer := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		answer += processA(scanner.Text(), priority)
	}

	fmt.Printf("The answer is %d\n", answer)
}
