package main

import (
	"fmt"
	"sort"
)

//Write a function that takes in a non-empty array of distinct integers and an
//integer representing a target sum. The function should find all triplets in
//the array that sum up to the target sum and return a two-dimensional array of
//all these triplets. The numbers in each triplet should be ordered in ascending
//order, and the triplets themselves should be ordered in ascending order with
//respect to the numbers they hold.

func main() {

	arr := []int{12,3,1,2,-6,5,-8,6}

	targetSum := 0

	fmt.Println("Three Number Sum: ", threeNumberSum(arr, targetSum))
}

func threeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	// Note distinct integers' means there is duplicates in the array
	// Sort Array
	sort.Ints(array)
	arrLen := len(array)
	currentIndex := 0

	threeSumArr := [][]int{}

	for currentIndex < arrLen - 2 {
		leftValue := currentIndex + 1
		rightValue := arrLen - 1
		for leftValue < rightValue {
			currentSum := array[currentIndex] + array[leftValue] + array[rightValue]
			if currentSum < target  {
				leftValue++
			}
			if currentSum > target  {
				rightValue--
			}
			if  currentSum == target {
				threeSumArr = append(threeSumArr, []int{array[currentIndex], array[leftValue], array[rightValue]})

				leftValue++
				rightValue--
			}
		}
		currentIndex++
	}
	return threeSumArr
}

