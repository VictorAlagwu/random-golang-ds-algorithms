package main

import (
	"fmt"
)


//Given an array of integers between 1 and n ,
//inclusive, where n is the length of the array, write a function
//that returns the first integer that appears more than once (when the array is
//read from left to right).
//
//
//In other words, out of all the integers that might occur more than once in the
//input array, your function should return the one whose first duplicate value
//has the minimum index.

//If no integer appears more than once, your function should return  -1.
//Note that you're allowed to mutate the input array.
//Sample Input #1
//array: [2, 1, 5, 2, 3, 3, 4]
//
//Sample Output #1 2 // 2 is the first integer that appears more than once.</span>

func main() {
	arr := []int{2, 1, 5, 2, 3, 3, 4}
	fmt.Println("Three Number Sum (2-pointer): ", firstDuplicateValue(arr))
}

func firstDuplicateValue(array []int) int {
	// Write your code here.
	checker := map[int]bool{}
	currentIndex := 0
	value := -1

	if len(array) <= 1 {
		return -1
	}

	for currentIndex < len(array) {
		currentValue := array[currentIndex]
		if _, found := checker[currentValue]; found {
			value = currentValue
			break
		}
		checker[currentValue] = true
		currentIndex++
	}
	return value
}
