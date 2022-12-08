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

	stacks := make([][]byte, 9)

	inputCollected := false
	processingMoves := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "1") > -1 {
			inputCollected = true
		}

		if strings.Index(line, "move") > -1 {
			processingMoves = true
		}

		if !inputCollected {
			stacks = getStacks(scanner.Text(), stacks)
		}

		if processingMoves {
			stacks = processMoves(line, stacks)
		}
	}

	fmt.Println("Top boxes: ")
	var answer []byte
	for _, x := range stacks {
		answer = append(answer, x[len(x)-1])
	}
	fmt.Println(string(answer))
}

func getStacks(line string, stacks [][]byte) [][]byte {
	i := 0
	j := 0
	for i < len(line) {
		value := line[i+1]

		stacks = insertAtFront(stacks, j, value)

		i += 4
		j += 1
	}

	return stacks
}

func processMoves(line string, stacks [][]byte) [][]byte {
	move := strings.Split(line, " ")
	numberToMove, _ := strconv.Atoi(move[1])
	from, _ := strconv.Atoi(move[3])
	to, _ := strconv.Atoi(move[5])

	stacks = moveCrate(stacks, numberToMove, from-1, to-1)
	return stacks
}

func insertAtFront(stacks [][]byte, stackIndex int, value byte) [][]byte {
	if value != ' ' {
		stacks[stackIndex] = append([]byte{value}, stacks[stackIndex]...)
	}

	fmt.Printf("%q\n", stacks)

	return stacks
}

func moveCrate(stacks [][]byte, numberToMove, fromStack, toStack int) [][]byte {
	// fmt.Printf("move %d from %d to %d", numberToMove, fromStack+1, toStack+1)
	i := 0
	for i < numberToMove {
		value := stacks[fromStack][len(stacks[fromStack])-1]
		stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
		stacks[toStack] = append(stacks[toStack], value)

		i++
	}

	return stacks
}
