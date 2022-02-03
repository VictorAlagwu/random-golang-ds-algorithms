package main

import "fmt"

func main() {

	nums := []int{2,7,11,15}
	target := 9
	fmt.Println("Two Sum II:", twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
