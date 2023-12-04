package main

import (
	"os"
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

/* First round */

func firstRound() {
	result := 0
	for _, line := range readInput("input.txt") {
		coordsFirst := 0
		coordsSecond := 0
		firstFound := false
		for _, character := range []rune(line) {
			if !unicode.IsDigit(character) {
				continue
			}
			if !firstFound {
				coordsFirst = int(character - '0')
				firstFound = true
			}
			coordsSecond = int(character - '0')
		}
		result += coordsFirst*10 + coordsSecond
	}
	println(result)
}

/* Second round */

var digits = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

type WordDigitsTrie struct {
	children  map[rune]*WordDigitsTrie
	wordValue int
}

func initTrie() (rootNode *WordDigitsTrie) {
	rootNode = &WordDigitsTrie{}
	for word, value := range digits {
		insertWord(rootNode, word, value)
	}
	return
}

func insertWord(rootNode *WordDigitsTrie, word string, value int) {
	current := rootNode
	for _, letter := range []rune(word) {
		if current.children == nil {
			current.children = make(map[rune]*WordDigitsTrie)
		}
		if current.children[letter] == nil {
			current.children[letter] = &WordDigitsTrie{}
		}
		current = current.children[letter]
	}
	current.wordValue = value
}

var trie = initTrie()

func secondRound() {
	result := 0
	for _, line := range readInput("input.txt") {
		first := 0
		second := 0
		currentTrie := trie
		runes := []rune(line)
		for idx := 0; idx < len(runes); idx++ {
			char := runes[idx]
			parsed := 0
			if unicode.IsDigit(char) {
				parsed = int(char - '0')
			} else {
				localIdx := idx
				for localIdx < len(runes) && currentTrie.children[runes[localIdx]] != nil {
					currentTrie = currentTrie.children[runes[localIdx]]
					if currentTrie.wordValue > 0 {
						parsed = currentTrie.wordValue
						break
					}
					localIdx++
				}
			}
			if parsed != 0 {
				if first == 0 {
					first = parsed
				}
				second = parsed
			}
			currentTrie = trie
		}
		result += first*10 + second
	}
	println(result)
}

func main() {
	firstRound()
	secondRound()
}
