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
