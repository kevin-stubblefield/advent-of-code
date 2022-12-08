package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")

	s := bufio.NewScanner(f)
	s.Scan()
	fmt.Println(getAnswer(s.Text()))
}

func getAnswer(input string) int {
Outer:
	for i, _ := range input {
		if i < 14 {
			continue
		}

		s := input[i-14 : i]
		set := make(map[rune]bool)
		for _, x := range s {
			if _, found := set[x]; found {
				continue Outer
			} else {
				set[x] = true
			}
		}

		return i
	}
	return -1
}
