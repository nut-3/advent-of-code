package main

import (
	"os"
	"strings"
)

var (
	result  = map[string]int{"A Z": 0, "B X": 0, "C Y": 0, "A X": 3, "B Y": 3, "C Z": 3, "A Y": 6, "B Z": 6, "C X": 6}
	rps     = map[string]int{"X": 1, "Y": 2, "Z": 3}
	result2 = map[string]int{"X": 0, "Y": 3, "Z": 6}
	rps2    = map[string]int{"A Z": 2, "B X": 1, "C Y": 3, "A X": 3, "B Y": 2, "C Z": 1, "A Y": 1, "B Z": 3, "C X": 2}
)

func calcSingleOutcome(line string) int {
	myMove := strings.Split(line, " ")[1]
	return result[line] + rps[myMove]
}

func firstRound() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}
	var result int
	for _, str := range strings.Split(string(file), "\n") {
		result += calcSingleOutcome(str)
	}
	println(result)
}

func calcSingleOutcome2(line string) int {
	result := strings.Split(line, " ")[1]
	return rps2[line] + result2[result]
}

func secondRound() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}
	result := 0
	for _, str := range strings.Split(string(file), "\n") {
		result += calcSingleOutcome2(str)
	}
	println(result)
}

func main() {
	firstRound()
	secondRound()
}
