package main

import "fmt"

//Write a function that takes in a non-empty array of integers and returns an
//array of the same length, where each element in the output array is equal to
//the product of every other number in the input array.
func main() {
	arr := []int{5, 1, 4, 2}
	fmt.Println("Array of Products: ", arrayOfProductsBruteForce(arr))
}

func arrayOfProductsBruteForce(array []int) []int {
	// Write your code here.
	arr := []int{}
	for i := 0; i < len(array); i++ {
		currentMultiple := 1
		for j := 0; j < len(array); j++ {
			if i == j {
				continue
			}
			currentMultiple *= array[j]
		}
		arr = append(arr, currentMultiple)
	}
	return arr
}