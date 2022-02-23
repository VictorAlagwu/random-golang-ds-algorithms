package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return al the triplets [nums[i], nums[j], num[k]]
// such that i != j, i != k, and j != k and nums[i] + nums[j] + nums[k] == 0.
// e.g Input = [-1,0,1,2,-1,-4] ... output = [[-1,-1,2],[-1,0,1]]
func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println("Three Sum:", threeSumI(nums))
}

func threeSumI(nums []int) [][]int {
	threeSumSlice := make(map[[3]int][]int)
	var threeSumArr [][]int
	// Brute Force Approach

	for i := 0; i < len(nums) - 2; i++ {
		for j := i + 1; j < len(nums) - 1; j++ {
			for k := j + 1; k < len(nums) ; k++ {
				arr := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(arr[:])
				if nums[i] + nums[j] + nums[k] == 0 {
					threeSumSlice[arr] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	for _, sum := range threeSumSlice {
		threeSumArr = append(threeSumArr, sum)
	}
	return threeSumArr
}

func threeSumII(nums []int) [][]int {
	//Firstly, sort the array
	sort.Ints(nums)
	var threeSumArr [][]int

	for idx1 := 0; idx1 < len(nums) - 2; idx1++ {
		if idx1 > 0 && nums[idx1] == nums[idx1 - 1] {
			continue
		}

		idx2 := idx1 + 1
		idx3 := len(nums) - 1
		for idx2 < idx3 {
			sum := nums[idx1] + nums[idx2] + nums[idx3]
			if sum == 0 {
				threeSumArr = append(threeSumArr, []int{nums[idx1], nums[idx2], nums[idx3]})
				idx3--

				for idx2 < idx3 && nums[idx3] == nums[idx3 + 1] {
					idx3--
				}
			} else if sum > 0 {
				idx3--
			} else {
				idx2++
			}
		}
	}

	return threeSumArr
}
