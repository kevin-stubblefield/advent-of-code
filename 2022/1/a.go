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
	max := 0
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

		if calories[currentElf] > calories[max] {
			max = currentElf
		}
	}

	fmt.Printf("there are %d elves\n", currentElf+1)
	fmt.Printf("The elf with the most calories is %d and is carrying %d calories \n", max+1, calories[max])

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
