package main

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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
	resultSync := new(atomic.Uint64)
	wg := new(sync.WaitGroup)
	defer func() {
		wg.Wait()
		result = int(resultSync.Load())
	}()
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			cards := strings.Split(strings.Split(line, ":")[1], "|")
			winningNums := strings.Fields(strings.TrimSpace(cards[0]))
			var numbers []int
			cardNumbers := strings.Fields(strings.TrimSpace(cards[1]))
			for _, numStr := range cardNumbers {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
			slices.Sort(numbers)
			countPoints := -1
			for _, winNum := range winningNums {
				winInt, _ := strconv.Atoi(winNum)
				_, found := slices.BinarySearch(numbers, winInt)
				if found {
					countPoints++
				}
			}
			if countPoints >= 0 {
				points := uint64(math.Pow(float64(2), float64(countPoints)))
				resultSync.Add(points)
			}
		}(line)
	}
	return
}

/* Second round */

func getCardDetails(line string) (cardNum int, winningNums []int, cardNums []int, err error) {
	cardParts := strings.Split(line, ":")
	cardNum, err = strconv.Atoi(strings.Fields(cardParts[0])[1])
	cards := strings.Split(cardParts[1], "|")
	winningChars := strings.Fields(strings.TrimSpace(cards[0]))
	cardNumbers := strings.Fields(strings.TrimSpace(cards[1]))
	for _, numStr := range cardNumbers {
		var num int
		num, err = strconv.Atoi(numStr)
		cardNums = append(cardNums, num)
	}
	slices.Sort(cardNums)
	for _, char := range winningChars {
		var num int
		num, err = strconv.Atoi(char)
		winningNums = append(winningNums, num)
	}
	return
}

func processCard(line string, wg *sync.WaitGroup, resultCardMap *sync.Map) {
	defer wg.Done()
	cardNum, winningNums, cardNums, err := getCardDetails(line)
	if err != nil {
		panic("Error getting card details")
	}
	countPoints := 0
	for _, winNum := range winningNums {
		_, found := slices.BinarySearch(cardNums, winNum)
		if found {
			countPoints++
		}
	}
	if countPoints > 0 {
		resultCardMap.Store(cardNum, countPoints)
	}
}

func secondRound(input string) (result int) {
	data := readInput(input)
	resultCardMap := new(sync.Map)
	wg := new(sync.WaitGroup)
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		wg.Add(1)
		go processCard(line, wg, resultCardMap)
	}
	wg.Wait()
	resultArray := make([]int, len(data))
	for cardNum := 1; cardNum <= len(data); cardNum++ {
		resultArray[cardNum-1]++
		value, ok := resultCardMap.Load(cardNum)
		if !ok {
			continue
		}
		nextNumbers := value.(int)
		for nextCardNum := cardNum; nextCardNum < len(data) && nextCardNum < cardNum+nextNumbers; nextCardNum++ {
			resultArray[nextCardNum] += resultArray[cardNum-1]
		}
	}
	for _, number := range resultArray {
		result += number
	}
	return
}

func main() {
	assert(firstRound, "example1.txt", 13)
	println(firstRound("input.txt"))
	assert(secondRound, "example2.txt", 30)
	println(secondRound("input.txt"))
}
