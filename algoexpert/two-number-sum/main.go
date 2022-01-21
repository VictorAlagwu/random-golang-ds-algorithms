package main

import (
	"fmt"
	//"sort"
)

func main() {
	arr := []int{1,-2,3,4,-4,6,-14,8,2}

	target := 4
	fmt.Println("Number Sum: O(n^2)", twoNumberSumSolutionOne(arr, target))
	fmt.Println("Number Sum: O(n)", twoNumberSumSolutionTwo(arr, target))
}

func twoNumberSumSolutionOne(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				continue;
			}
			if array[i] + array[j] == target {
				return []int{array[i],  array[j]}
			}
		}
	}

	return []int{}
}

func twoNumberSumSolutionTwo(array []int, target int) []int {
	nums := map[int]bool{}

	for _, num := range array {
		checkValue := target - num
		if _, found := nums[checkValue]; found {
			return []int{checkValue, num}
		}
		nums[num] = true
	}
	return []int{}
}