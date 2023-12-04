package main

import (
	"os"
	"strconv"
	"strings"
)

var monkeys = make(map[string]*Monkey)

type Monkey struct {
	first   string
	second  string
	num     int
	operand string
}

func (monkey *Monkey) Calculate() int {
	if len(monkey.first) == 0 {
		return monkey.num
	}
	first := monkeys[monkey.first].Calculate()
	second := monkeys[monkey.second].Calculate()
	switch monkey.operand {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	}
	return 0
}

func readInput() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Файл не найден")
	}
	return string(file)
}

func firstRound() {
	for _, str := range strings.Split(readInput(), "\n") {
		parsedString := strings.Split(str, ":")
		name := parsedString[0]
		data := strings.TrimSpace(parsedString[1])
		num, numErr := strconv.Atoi(data)
		if numErr != nil {
			dataStrings := strings.Split(data, " ")
			monkeys[name] = &Monkey{first: dataStrings[0], second: dataStrings[2], operand: dataStrings[1]}
		} else {
			monkeys[name] = &Monkey{num: num}
		}
	}
	println(monkeys["root"].Calculate())
}

func main() {
	firstRound()
}
