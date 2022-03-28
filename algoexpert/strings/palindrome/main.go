package main

import "fmt"

//Write a function that takes in a non-empty string and that returns a boolean representing whether the string is a palindrome.
//	A palindrome is defined as a string that's written the same forward and backward.
//Note that single-character strings are palindromes. Sample Input string = "abcdcba"
//Sample Output true // it's written the same forward and backward

func main() {

	order := "abcdcba"
	fmt.Println("Is Palindrome: ", isPalindrome(order))
}

func isPalindrome(str string) bool {
	// Write your code here.
	leftIndex := 0
	rightIndex := len(str) -1

	for leftIndex <= rightIndex {
		if str[leftIndex] != str[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}