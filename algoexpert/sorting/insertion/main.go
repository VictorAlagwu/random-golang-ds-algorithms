package main

import "fmt"

func main() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println("Insertion Sort: ", insertionSort(arr))
}

func insertionSort(array []int) []int {
	// Write your code here.
	for idx,_ := range array {
		j := idx
		for j > 0 && array[j] < array[j - 1] {
			array[j], array[j - 1] = array[j - 1], array[j]
			j -= 1
		}
	}
	return array
}
