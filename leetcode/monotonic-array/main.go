package main
//https://leetcode.com/problems/monotonic-array
func IsMonotonic(nums []int) bool {
	// Write your code here.
	if len(nums) <= 2 {
		return true
	}
	status := false
	currentActive := ""
	for i := 0; i < len(nums); i++ {
		if i <= 0 || i == len(nums) - 1{
			continue
		}
		currentValue := nums[i]
		nextValue := nums[i + 1]
		previousValue := nums[i-1]

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