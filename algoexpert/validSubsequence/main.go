package main

import "fmt"

func main() {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println("Number Sum: O(n)", isValidSubsequence(arr, sequence))
}

func isValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	arrayIdx := 0
	sequenceIdx := 0

	for  arrayIdx < len(array) && sequenceIdx < len(sequence) {
		if array[arrayIdx] == sequence[sequenceIdx] {
			sequenceIdx ++
		}
		arrayIdx++
	}
	
	if len(sequence) == sequenceIdx {
		return true
	} else {
		return false
	}
}
