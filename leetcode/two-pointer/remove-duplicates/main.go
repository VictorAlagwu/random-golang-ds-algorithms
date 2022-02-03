package main

import "fmt"

func main() {
	nums := []int{1,1,1,2,2,3,3,4,4,5,5,6,7}
	fmt.Println("Remove duplicates:", removeDuplicates(nums))

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	lastUniqueIndex := 0

	for i < len(nums) {
		if nums[lastUniqueIndex] != nums[i] {
			lastUniqueIndex++
			nums[lastUniqueIndex] = nums[i]
		}
		i++
	}
	return lastUniqueIndex + 1
}