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
	pointMap["ROCK"] = 1
	pointMap["PAPER"] = 2
	pointMap["SCISSORS"] = 3

	pointMap["X"] = 0
	pointMap["Y"] = 3
	pointMap["Z"] = 6

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
	desiredOutcome := moves[1]

	myMove := play(desiredOutcome, opponentMove)

	points := pointMap[myMove] + pointMap[desiredOutcome]

	return points
}

func play(desiredOutcome, opponentMove string) string {
	switch opponentMove + desiredOutcome {
	case "AX":
		return "SCISSORS"
	case "AY":
		return "ROCK"
	case "AZ":
		return "PAPER"
	case "BX":
		return "ROCK"
	case "BY":
		return "PAPER"
	case "BZ":
		return "SCISSORS"
	case "CX":
		return "PAPER"
	case "CY":
		return "SCISSORS"
	case "CZ":
		return "ROCK"
	default:
		log.Fatalln("WTF")
		return "WTF"
	}
}
