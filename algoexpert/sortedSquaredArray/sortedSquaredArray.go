package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 5, 6, 8, 9}
	fmt.Println("Sorted Squared Array: O(nlogn) ", sortedSquaredArrayBruteForeApproach(arr))
	fmt.Println("Sorted Squared Array: O(n)", sortedSquaredArraySolutionTwo(arr))
}

func sortedSquaredArrayBruteForeApproach(array []int) []int {
	// Write your code here
	arr := make([]int, len(array))

	for idx, value := range array {
		arr[idx] = value * value
	}
	sort.Ints(arr)
	return arr
}

func sortedSquaredArraySolutionTwo(array []int) []int {
	startIdx := 0
	endIdx := len(array) - 1
	currentIndex := 0
	arr := make([]int, len(array))

	for currentIndex < len(array) {
		currentIndex += 1
		arrayStartValue := array[startIdx] * array[startIdx]
		arrayEndValue := array[endIdx] * array[endIdx]
		if  arrayStartValue < arrayEndValue {
			arr[len(array) - currentIndex] = arrayEndValue
			endIdx -= 1
		} else {
			arr[len(array) - currentIndex] = arrayStartValue
			startIdx += 1
		}
	}
	return arr
}
