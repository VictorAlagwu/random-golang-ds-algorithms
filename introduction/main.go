package main

import (
	"fmt"
)

func main(){
	data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	nums := [][]int{data}
	fmt.Println("Average: ", findAverage(data))
	fmt.Println("Sum of multiple array: ", sumOfTwoDim(nums))
	fmt.Println("Largest element: ", largestElement(data))
	fmt.Println("Smallest element: ", smallestElement(data))
	fmt.Println("Second Largest number: ", secondLargestElement(data))
	fmt.Println("Sum of number: ", sum(6))

}


func findAverage(s []int) int {
	var totalValue int = 0;
	for _, i := range s {
		totalValue += i
	}
	return totalValue / len(s)
}

func sumOfTwoDim(nums [][]int) int {
	var sum int = 0;
	for _, num := range nums {
		for _, i := range num {
			sum += i
		}
	}
	return sum
}

func largestElement(s []int)int {
	var maxValue int = 0;
	for _, i := range s {
		if i > maxValue {
			maxValue = i
		}
	}
	return maxValue
}

func smallestElement(s []int)int {
	var minValue int = 0;
	for _, i := range s {
		if i < minValue {
			minValue = i
		}
	}
	return minValue
}

func secondLargestElement(s []int) int {
	var firstMaxValue int = 0;
	var secondMaxValue int = 0;

	for _, i := range s {
		if i > firstMaxValue {
			firstMaxValue = i
		}
	}

	for _, i := range s {
		if i > secondMaxValue && i != firstMaxValue {
			secondMaxValue = i
		}
	}
	return secondMaxValue
}

func sum(n int) int {
	var totalValue int = 0;
	for i := 0; i < n; i++ {
		totalValue += i
	}
	return totalValue
}

