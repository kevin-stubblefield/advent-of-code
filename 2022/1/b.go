package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentElf := 0
	elfWithMost := 0
	elfWithSecondMost := 0
	elfWithThirdMost := 0
	var calories []int
	calories = append(calories, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// find the sum of each elf's calories
		if scanner.Text() == "" {
			currentElf++
			calories = append(calories, 0)
		} else {
			currentCalories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			calories[currentElf] += currentCalories
		}

		if calories[currentElf] > calories[elfWithMost] {
			elfWithThirdMost = elfWithSecondMost
			elfWithSecondMost = elfWithMost
			elfWithMost = currentElf
		} else if calories[currentElf] > calories[elfWithSecondMost] {
			elfWithThirdMost = elfWithSecondMost
			elfWithSecondMost = currentElf
		} else if calories[currentElf] > calories[elfWithThirdMost] {
			elfWithThirdMost = currentElf
		}
	}

	fmt.Printf("there are %d elves\n", currentElf+1)
	fmt.Printf("The elves with the most calories are %d, %d, %d and is carrying %d calories \n", elfWithMost+1, elfWithSecondMost+1, elfWithThirdMost+1, calories[elfWithMost]+calories[elfWithSecondMost]+calories[elfWithThirdMost])

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
