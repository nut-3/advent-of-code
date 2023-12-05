package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
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

var coordsShift = [][]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}

type matrix [][]rune

/* First round */

func isSymbol(char *rune) bool {
	return !unicode.IsDigit(*char) && '.' != *char
}

func (m matrix) findNumbers(y int, x int) (result int) {
	for _, shifts := range coordsShift {
		curY := y + shifts[0]
		if curY < 0 || curY >= len(m) {
			continue
		}
		curX := x + shifts[1]
		if !unicode.IsDigit(m[curY][curX]) {
			continue
		}
		for curX < len(m[curY]) && unicode.IsDigit(m[curY][curX]) {
			curX++
		}
		curX--
		number := 0
		for digit := float64(0); curX >= 0 && unicode.IsDigit(m[curY][curX]); digit++ {
			number += int(m[curY][curX]-'0') * int(math.Pow(10, digit))
			m[curY][curX] = '.'
			curX--
		}
		result += number
	}
	return
}

func firstRound(input string) (result int) {
	data := readInput(input)
	mtx := make(matrix, len(data))
	for num, line := range data {
		mtx[num] = []rune(line)
	}
	for y := 0; y < len(mtx); y++ {
		for x := 0; x < len(mtx[y]); x++ {
			if isSymbol(&mtx[y][x]) {
				result += mtx.findNumbers(y, x)
			}
		}
	}
	return
}

/* Second round */

func isGear(char *rune) bool {
	return *char == '*'
}

func (m matrix) findGearNumbers(y int, x int) (result int) {
	var numbers []int
	for _, shifts := range coordsShift {
		curY := y + shifts[0]
		if curY < 0 || curY >= len(m) {
			continue
		}
		curX := x + shifts[1]
		if !unicode.IsDigit(m[curY][curX]) {
			continue
		}
		for curX < len(m[curY]) && unicode.IsDigit(m[curY][curX]) {
			curX++
		}
		curX--
		number := 0
		for digit := float64(0); curX >= 0 && unicode.IsDigit(m[curY][curX]); digit++ {
			number += int(m[curY][curX]-'0') * int(math.Pow(10, digit))
			m[curY][curX] = '.'
			curX--
		}
		numbers = append(numbers, number)
	}
	if len(numbers) == 2 {
		result = numbers[0] * numbers[1]
	}
	return
}

func secondRound(input string) (result int) {
	data := readInput(input)
	mtx := make(matrix, len(data))
	for num, line := range data {
		mtx[num] = []rune(line)
	}
	for y := 0; y < len(mtx); y++ {
		for x := 0; x < len(mtx[y]); x++ {
			if isGear(&mtx[y][x]) {
				result += mtx.findGearNumbers(y, x)
			}
		}
	}
	return
}

func main() {
	assert(firstRound, "example1.txt", 4361)
	println(firstRound("input.txt"))
	assert(secondRound, "example2.txt", 467835)
	println(secondRound("input.txt"))
}
