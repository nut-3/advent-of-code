package main

import (
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("File not found")
	}
	return strings.Split(string(file), "\n")
}

func assert(f func(string) int, testSet string, result int) {
	if f(testSet) != result {
		panic("Wrong answer!")
	}
}

/* First round */

var maxGemsMap = map[string]int{"red": 12, "green": 13, "blue": 14}

func firstRound(input string) (result int) {
	for _, line := range readInput(input) {
		game := strings.Split(line, ": ")
		gameNum, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		setsOfGems := strings.Split(game[1], "; ")
		includeSetStats := true
		for _, gemsSet := range setsOfGems {
			gemsKinds := strings.Split(gemsSet, ", ")
			for _, gemsKind := range gemsKinds {
				gemsStat := strings.Split(gemsKind, " ")
				gemsNum, _ := strconv.Atoi(gemsStat[0])
				if maxGemsMap[gemsStat[1]] < gemsNum {
					includeSetStats = false
					break
				}
			}
			if !includeSetStats {
				break
			}
		}
		if !includeSetStats {
			continue
		}
		result += gameNum
	}
	return
}

/* Second round */

func secondRound(input string) (result int) {
	for _, line := range readInput(input) {
		game := strings.Split(line, ": ")
		setsOfGems := strings.Split(game[1], "; ")
		currentGemsMinAmount := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, gemsSet := range setsOfGems {
			gemsKinds := strings.Split(gemsSet, ", ")
			for _, gemsKind := range gemsKinds {
				gemsStat := strings.Split(gemsKind, " ")
				gemsNum, _ := strconv.Atoi(gemsStat[0])
				if currentGemsMinAmount[gemsStat[1]] < gemsNum {
					currentGemsMinAmount[gemsStat[1]] = gemsNum
				}
			}
		}
		result += currentGemsMinAmount["red"] * currentGemsMinAmount["green"] * currentGemsMinAmount["blue"]
	}
	return
}

func main() {
	assert(firstRound, "example1.txt", 8)
	println(firstRound("input.txt"))
	assert(secondRound, "example2.txt", 2286)
	println(secondRound("input.txt"))
}
