package main

import "fmt"

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println("Two Sum II:", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	arr := make(map[int]int)
	currentIndex := 0
	for currentIndex < len(nums) {
		v := nums[currentIndex]
		s := target - v

		if key, value := arr[s]; value {
			return []int{key, currentIndex}
		}
		arr[v] = currentIndex
		currentIndex++

	}
	return nil
}