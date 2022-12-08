package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Could not open file")
	}

	pointMap := make(map[string]int)
	pointMap["X"] = 1
	pointMap["Y"] = 2
	pointMap["Z"] = 3

	pointMap["AX"] = 3
	pointMap["AY"] = 6
	pointMap["AZ"] = 0
	pointMap["BY"] = 3
	pointMap["BZ"] = 6
	pointMap["BX"] = 0
	pointMap["CZ"] = 3
	pointMap["CX"] = 6
	pointMap["CY"] = 0

	var points int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points += calculatePointsForLine(scanner.Text(), pointMap)
	}

	fmt.Printf("You scored %d points!", points)
}

func calculatePointsForLine(line string, pointMap map[string]int) int {
	moves := strings.Split(line, " ")
	opponentMove := moves[0]
	myMove := moves[1]

	points := pointMap[myMove] + pointMap[opponentMove+myMove]

	return points
}
