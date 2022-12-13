package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Monkey struct {
	items          []int
	operator       string
	operand        int
	test           int
	throwsToTrue   int
	throwsToFalse  int
	itemsInspected int
}

func NewMonkey(itemStrings []string, operator string, test, operand, throwsToTrue, throwsToFalse int) *Monkey {
	var items []int
	for _, i := range itemStrings {
		n, _ := strconv.Atoi(i)
		items = append(items, n)
	}

	return &Monkey{
		items:         items,
		operator:      operator,
		operand:       operand,
		test:          test,
		throwsToTrue:  throwsToTrue,
		throwsToFalse: throwsToFalse,
	}
}

func (m *Monkey) takeTurn(monkeys []*Monkey, manageWorryLevels bool, reductionValue int) {
	for _, i := range m.items {
		m.itemsInspected++

		if !manageWorryLevels {
			i %= reductionValue
		}

		var realOperand int
		if m.operand == -1 {
			realOperand = i
		} else {
			realOperand = m.operand
		}

		switch m.operator {
		case "+":
			i += realOperand
		case "*":
			i *= realOperand
		}

		if manageWorryLevels {
			i = int(math.Floor(float64(i / 3)))
		}

		var throwTo int
		if i%m.test == 0 {
			throwTo = m.throwsToTrue
		} else {
			throwTo = m.throwsToFalse
		}
		monkeys[throwTo].items = append(monkeys[throwTo].items, i)

		m.items = m.items[1:]
	}
}

func part1(monkeys []*Monkey) {
	round := 1
	for round <= 20 {
		for _, m := range monkeys {
			m.takeTurn(monkeys, true, 0)
		}

		// for i, m := range monkeys {
		// 	fmt.Printf("Monkey #%d, itemsInspected: %d, items: %v\n", i, m.itemsInspected, m.items)
		// }
		round++
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected < monkeys[j].itemsInspected
	})

	fmt.Printf("Answer: %d\n", monkeys[len(monkeys)-1].itemsInspected*monkeys[len(monkeys)-2].itemsInspected)
}

func part2(monkeys []*Monkey) {
	reductionValue := 1
	for _, m := range monkeys {
		reductionValue *= m.test
	}

	round := 1
	for round <= 10000 {
		for _, m := range monkeys {
			m.takeTurn(monkeys, false, reductionValue)
		}

		// if round <= 20 {
		// 	fmt.Printf("Round #%d\n", round)
		// 	for i, m := range monkeys {
		// 		fmt.Printf("Monkey #%d, itemsInspected: %d, items: %v\n", i, m.itemsInspected, m.items)
		// 	}
		// }
		round++
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected < monkeys[j].itemsInspected
	})
	// for i, m := range monkeys {
	// 	fmt.Printf("Monkey #%d, itemsInspected: %d, items: %v\n", i, m.itemsInspected, m.items)
	// }

	fmt.Printf("Answer: %d\n", monkeys[len(monkeys)-1].itemsInspected*monkeys[len(monkeys)-2].itemsInspected)
}

func createMonkeys(f string) []*Monkey {
	monkeyStrings := strings.Split(f, "\r\n\r\n")

	var items []string
	var operator string
	var test, throwsToTrue, throwsToFalse int

	var monkeys []*Monkey

	for _, m := range monkeyStrings {
		parts := strings.Split(m, "\r\n")
		items = strings.Split(strings.TrimPrefix(parts[1], "  Starting items: "), ", ")
		operation := strings.Split(strings.TrimPrefix(parts[2], "  Operation: new = old "), " ")
		operator = operation[0]
		operand, err := strconv.Atoi(operation[1])
		if err != nil {
			operand = -1
		}
		test, _ = strconv.Atoi(strings.TrimPrefix(parts[3], "  Test: divisible by "))
		throwsToTrue, _ = strconv.Atoi(strings.TrimPrefix(parts[4], "    If true: throw to monkey "))
		throwsToFalse, _ = strconv.Atoi(strings.TrimPrefix(parts[5], "    If false: throw to monkey "))

		monkey := NewMonkey(items, operator, test, operand, throwsToTrue, throwsToFalse)
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func main() {
	f, _ := os.ReadFile("input.txt")

	monkeys := createMonkeys(string(f))

	start := time.Now()
	part1(monkeys)
	fmt.Printf("Part 1: %s\n", time.Since(start))

	monkeys = createMonkeys(string(f))
	start = time.Now()
	part2(monkeys)
	fmt.Printf("Part 2: %s\n", time.Since(start))
}
