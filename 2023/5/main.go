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
	testResult := f(testSet)
	if testResult != result {
		panic("Wrong answer: " + strconv.FormatInt(int64(testResult), 10))
	}
}

/* First round */

func firstRound(input string) (result int) {
	data := readInput(input)
	return
}

/* Second round */

func secondRound(input string) (result int) {
	data := readInput(input)
	return
}

func main() {
	assert(firstRound, "example1.txt", 13)
	println(firstRound("input.txt"))
	assert(secondRound, "example2.txt", 30)
	println(secondRound("input.txt"))
}
