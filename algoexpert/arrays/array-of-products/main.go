package main

import "fmt"

//Write a function that takes in a non-empty array of integers and returns an
//array of the same length, where each element in the output array is equal to
//the product of every other number in the input array.
func main() {
	arr := []int{5, 1, 4, 2}
	fmt.Println("Array of Products (Brute Force): ", arrayOfProductsBruteForce(arr))
	fmt.Println("Array of Products (Two Pointer): ", arrayOfProducts(arr))
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

func arrayOfProducts(array []int) []int {
	lengthOfArr  := len(array)
	currentIndex := 0
	leftIndex := 0
	rightIndex := lengthOfArr - 1
	currentLeftProduct := 1
	currentRightProduct := 1
	leftProducts := make([]int, lengthOfArr)
	rightProducts := make([]int, lengthOfArr)
	products := make([]int, lengthOfArr)
	for i := range leftProducts {
		leftProducts[i] = 1
		rightProducts[i] = 1
		products[i] = 1
	}

	for leftIndex < lengthOfArr {
		leftProducts[leftIndex] = currentLeftProduct
		rightProducts[rightIndex] = currentRightProduct

		currentLeftProduct *= array[leftIndex]
		currentRightProduct *= array[rightIndex]

		leftIndex++
		rightIndex--
	}
	for currentIndex < len(leftProducts) {
		products[currentIndex] = leftProducts[currentIndex] * rightProducts[currentIndex]
		currentIndex++
	}

	return products
}
