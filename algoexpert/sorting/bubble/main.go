package main

import "fmt"

func main() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println("Bubble Sort: ", bubbleSort(arr))
}

func bubbleSort(array []int) []int {
	// Write your code here.
	currentIndex := 0
	lengthOfArray := len(array)

	for currentIndex < lengthOfArray {
		leftIndex := currentIndex
		rightIndex := currentIndex + 1
		if rightIndex == lengthOfArray {
			lengthOfArray--
			currentIndex = 0
			continue
		}
		if array[leftIndex] > array[rightIndex] {
			array[leftIndex], array[rightIndex]  = array[rightIndex], array[leftIndex]
		}

		currentIndex++
	}
	return array
}
