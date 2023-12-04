package main

import (
	"os"
	"strings"
)

func letterIndex(letter rune) int {
	if letter < 'a' {
		return int(letter-'A') + 26
	}
	return int(letter - 'a')
}

func processRucksack(rs string) int {
	middle := len(rs) / 2
	first := [26 * 2]int{}
	runes := []rune(rs)
	for i := 0; i < len(rs); i++ {
		runeIdx := letterIndex(runes[i])
		if i < middle {
			first[runeIdx]++
		} else {
			if first[runeIdx] > 0 {
				return runeIdx + 1
			}
		}
	}
	return 0
}

func firstRound() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}

	result := 0
	for _, str := range strings.Split(string(file), "\n") {
		priority := processRucksack(str)
		result += priority
	}

	println(result)
}

func secondRun() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}

	result := 0
	setOfSacks := [3]string{}
	counter := 0
	for _, str := range strings.Split(string(file), "\n") {
		setOfSacks[counter] = str
		counter++
		if counter > 2 {
			counter = 0
			priority := findCommonItemPriority(setOfSacks)
			result += priority
		}
	}

	println(result)
}

func findCommonItemPriority(sacks [3]string) int {
	first := [26 * 2]int{}
	second := [26 * 2]int{}
	runes := []rune(sacks[0])
	for _, item := range runes {
		runeIdx := letterIndex(item)
		first[runeIdx]++
	}
	runes = []rune(sacks[1])
	for _, item := range runes {
		runeIdx := letterIndex(item)
		second[runeIdx]++
	}
	runes = []rune(sacks[2])
	for _, item := range runes {
		runeIdx := letterIndex(item)
		if first[runeIdx] > 0 && second[runeIdx] > 0 {
			return runeIdx + 1
		}
	}
	return 0
}

func main() {
	firstRound()
	secondRun()
}
