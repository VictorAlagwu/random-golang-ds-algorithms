package main

import (
	"fmt"
	"math"
	"sort"
)

//Write a function that takes in two non-empty arrays of integers, finds the
//pair of numbers (one from each array) whose absolute difference is closest to
//zero, and returns an array containing these two numbers, with the number from
//the first array in the first position.


//Note that the absolute difference of two integers is the distance between
//them on the real number line. For example, the absolute difference of -5 and 5
//is 10, and the absolute difference of -5 and -4 is 1.


//You can assume that there will only be one pair of numbers with the smallest
//difference.
func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}
	fmt.Println("Smallest Difference:", smallestDifferenceII(array1, array2))

}


func smallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	arr := make([]int, 2)
	smallestDifference := 0
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			getAbs := array1[i] - array2[j];
			abs := math.Abs(float64(getAbs))
			if i == 0 && j == 0 {
				smallestDifference = int(abs)
				arr[0] = array1[i]
				arr[1] = array2[j]
			}else if int(abs) == smallestDifference {
				smallestDifference = int(abs)
				arr[0] = array1[i]
				arr[1] = array2[j]

			}else if int(abs) < smallestDifference {
				smallestDifference = int(abs)
				arr[0] = array1[i]
				arr[1] = array2[j]
			}else{

			}
		}
	}
	return arr
}

func smallestDifferenceII(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	arr := make([]int, 2)
	smallestDifference := math.MaxInt32
	currentArrayOneIndex := 0
	currentArrayTwoIndex := 0
	getAbs := math.MaxInt32
		for currentArrayOneIndex < len(array1) && currentArrayTwoIndex < len(array2) {

			firstValue := array1[currentArrayOneIndex]
			secondValue := array2[currentArrayTwoIndex]
			fmt.Println("First Value", firstValue)
			fmt.Println("Second Value", secondValue)
			fmt.Println("Smallest Value", smallestDifference)
			if  firstValue < secondValue {
				getAbs = secondValue - firstValue
				currentArrayOneIndex++
			} else if secondValue < firstValue  {
				getAbs = firstValue - secondValue
				currentArrayTwoIndex++
			}else {
				arr[0] = firstValue
				arr[1] = secondValue
				return arr
			}

			if smallestDifference > getAbs {
				smallestDifference = getAbs
				arr = []int{firstValue,  secondValue}
			}

		}
	return arr
}

