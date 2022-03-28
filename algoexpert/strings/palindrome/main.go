package main

import (
	"fmt"
	"strings"
)

//Write a function that takes in a non-empty string and that returns a boolean representing whether the string is a palindrome.
//	A palindrome is defined as a string that's written the same forward and backward.
//Note that single-character strings are palindromes. Sample Input string = "abcdcba"
//Sample Output true // it's written the same forward and backward

func main() {

	str := "abcdcba"
	sentence := "race a car"
	fmt.Println("Is Palindrome: ", isPalindrome(str))
	fmt.Println("Is Palindrome: ", isPalindromeForSentence(sentence))
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

func isPalindromeForSentence(str string) bool {
	// Write your code here.
	processedString := strings.ToLower(str)
	leftIndex := 0
	rightIndex := len(processedString) -1

	for leftIndex <= rightIndex {
		if !checkAlphaNumeric(processedString[leftIndex]) {
			leftIndex++
			continue
		}
		if ! checkAlphaNumeric(processedString[rightIndex]) {
			rightIndex--
			continue
		}
		if processedString[leftIndex] != processedString[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}

func checkAlphaNumeric(char uint8) bool {
	//a-z,A-Z,0-9
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}