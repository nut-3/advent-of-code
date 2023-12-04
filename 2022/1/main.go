package main

import (
	"container/heap"
	"os"
	"strconv"
	"strings"
)

func countMaxCalories() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}
	calories := 0
	maxCalories := 0
	for _, str := range strings.Split(string(file), "\n") {
		curCalories, strConvErr := strconv.Atoi(str)
		if strConvErr != nil {
			if maxCalories < calories {
				maxCalories = calories
			}
			calories = 0
		} else {
			calories += curCalories
		}
	}
	println(maxCalories)
}

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *IntHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func countMAxThreeSumCalories() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}
	calories := 0

	priorityQueue := make(IntHeap, 0)
	heap.Init(&priorityQueue)

	for _, str := range strings.Split(string(file), "\n") {
		curCalories, strConvErr := strconv.Atoi(str)
		if strConvErr != nil {
			heap.Push(&priorityQueue, calories)
			if priorityQueue.Len() > 3 {
				heap.Pop(&priorityQueue)
			}
			calories = 0
		} else {
			calories += curCalories
		}
	}
	maxSumCalories := 0
	for _, oneOfThreeMaxCalories := range priorityQueue {
		maxSumCalories += oneOfThreeMaxCalories
	}
	println(maxSumCalories)
}

func main() {
	countMaxCalories()
	countMAxThreeSumCalories()
}
