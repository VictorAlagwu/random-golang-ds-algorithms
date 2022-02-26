package main

import "fmt"

//Write a function that takes in an array of integers and returns a boolean
//representing whether the array is monotonic.


//An array is said to be monotonic if its elements, from left to right, are
//entirely non-increasing or entirely non-decreasing.

//Non-increasing elements aren't necessarily exclusively decreasing; they simply
//don't increase. Similarly, non-decreasing elements aren't necessarily
//exclusively increasing; they simply don't decrease.

//Note that empty arrays and arrays of one element are monotonic.
//Sample Input  [-1, -5, -10, -1100, -1100, -1101, -1102, -9001]
func main() {
	arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println("Three Number Sum (2-pointer): ", isMonotonic(arr))
}

func isMonotonic(array []int) bool {
	// Write your code here.
	if len(array) <= 2 {
		return true
	}
	status := false
	currentActive := ""
	for i := 0; i < len(array); i++ {
		if i <= 0 || i == len(array) - 1{
			continue
		}
		currentValue := array[i]
		nextValue := array[i + 1]
		previousValue := array[i-1]

		if currentValue == nextValue {
			status = true
			continue
		}

		if currentValue >= nextValue && currentValue <= previousValue {
			if currentActive == "increasing" {
				status = false
				break
			}
			status = true
			currentActive = "decreasing"
		} else if currentValue <= nextValue && currentValue >= previousValue {
			if currentActive == "decreasing" {
				status = false
				break
			}
			status = true
			currentActive = "increasing"
		} else {
			status = false
			break
		}
	}
	return status
}

