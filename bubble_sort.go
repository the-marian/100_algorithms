package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	someList := getIntRange(0, 10)
	shuffle(someList)
	fmt.Println("Before bubble sort:\n", someList)
	bubbleSort(someList)
	fmt.Println("After bubble sort:\n", someList)
}

func bubbleSort(s []int) []int {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
	return s
}

func getIntRange(from, to int) []int {
	sliceRange := make([]int, to-from)
	for i := range sliceRange {
		sliceRange[i] = i + 1
	}
	return sliceRange
}

func shuffle(slice []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}
