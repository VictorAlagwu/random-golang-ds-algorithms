package main

import "fmt"


//You're given an array of integers and an integer. Write a function that moves
//all instances of that integer in the array to the end of the array and returns
//the array.

//The function should perform this in place (i.e., it should mutate the input
//array) and doesn't need to maintain the order of the other integers.

func main() {

	nums := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	fmt.Println("Move Element to End:", moveElementToEnd(nums, toMove))
	fmt.Println("Move Element to End (Using 2 Pointers):", moveElementToEndII(nums, toMove))
}

func moveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	arr := []int{}
	arr2 := []int{}
	for i := 0; i < len(array); i++ {
		if toMove != array[i] {
			arr = append(arr, array[i])
		} else {
			arr2 = append(arr2, array[i])
		}
	}
	//Merge the two arrays...with the other
	return append(arr, arr2...)
}

func moveElementToEndII(array []int, toMove int) []int  {
	//Using two-pointers algorithm
	firstIndexValue := 0
	secondIndexValue := len(array) - 1

	for firstIndexValue < secondIndexValue {
		firstValue := array[firstIndexValue]
		secondValue := array[secondIndexValue]
		if secondValue == toMove {
			secondIndexValue--
		}

		if firstValue == toMove && secondValue != toMove {
			array[firstIndexValue] = secondValue
			array[secondIndexValue] = firstValue
			firstIndexValue++
			secondIndexValue--
		}

		if firstValue != toMove && secondValue != toMove {
			firstIndexValue++
		}
	}
	return array
}