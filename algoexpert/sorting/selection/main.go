package main

import "fmt"

func main() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println("Selection Sort: ", selectionSort(arr))
}

func selectionSort(array []int) []int {
	// Write your code here.
	smallestNumber := 0
	smallestNumberIndex := 0
	currentIndex := 0
	for currentIndex < len(array) {
		smallestNumber = array[currentIndex]
		smallestNumberIndex = currentIndex

		for j := currentIndex + 1; j < len(array); j++ {
			if array[j] < smallestNumber {
				smallestNumber = array[j]
				smallestNumberIndex = j
			}
		}
		if array[currentIndex] > smallestNumber {
			array[smallestNumberIndex], array[currentIndex] = array[currentIndex], array[smallestNumberIndex]
		}
		currentIndex++
	}
	return array
}
