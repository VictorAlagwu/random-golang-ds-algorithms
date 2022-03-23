package main

import "fmt"

//
//The Fibonacci sequence is defined as follows: the first number of the sequence
//is 0, the second number is 1, and the nth number is the sum of the (n - 1)th
//and (n - 2)th numbers. Write a function that takes in an integer
//n and returns the nth Fibonacci number.

//Important note: the Fibonacci sequence is often defined with its first two
//numbers as F0 = 0 and F1 = 1. For the purpose of this question,
// the first Fibonacci number is F0; therefore,
//getNthFib(1) is equal to F0, getNthFib(2)
//is equal to F1, etc..

func main() {
	value := 5
	fmt.Println("nth Fibonacci (Recursion): ", getNthFib(value))
}

func getNthFib(n int) int {
	// Write your code here.
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 0
	}
	fmt.Println("Current N", n)
	firstFib := + getNthFib(n - 1)
	secondFib := getNthFib(n - 2)
	fmt.Println("First N", firstFib)
	fmt.Println("Second N", secondFib)
	return firstFib + secondFib
}

func getNthFibUsingHashTable(n int) int {
	//Complete later
	return -1
}

