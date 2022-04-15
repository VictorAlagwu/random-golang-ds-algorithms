package main

import "fmt"

func main() {
	num := 5
	arr := []int{1, 2, 3, 4,2}
	str := "MADDAM"
	fmt.Println("Print Name: ", printName(num))
	fmt.Println("Print Linearly: ", printLinearly(1, num))
	fmt.Println("Print From N to One: ", printFromNToOne( num))
	fmt.Println("Print From One to N Using Backtracking: ", printLinearlyFromOneToN(num, num))
	fmt.Println("Print From N to One Using Backtracking: ", printFromNToOneBacktrack(1, num))
	fmt.Println("Summation of Number: ", summationOfNumber(num))
	fmt.Println("Reverse Array using Recursion using Two Pointer: ",
		reverseArrUsingRecursionUsingTwoPointer(0, len(arr) - 1,arr),
	)
	fmt.Println("Check if it's a palindrome: ",checkIfPalindromeUsingRecursion(0, str))
	fmt.Println("Fibonacci (Multiple Recursion call): ",fibonnaci(6))
}

// Print Name n times
func printName(n int) string {
	if n > 0 {
		fmt.Printf("Name %d\n", n)
		return printName(n - 1)
	}
	return ""
}

//Print Linearly from 1 to N
func printLinearly(start int, n int) string {
	if start > n {
		return ""
	}

	fmt.Printf("Name %d\n", start)
	return printLinearly(start + 1, n)
}

// Print from N to 1
func printFromNToOne(n int) string {
	if n < 1 {
		return ""
	}

	fmt.Printf("Name %d\n", n)
	return printFromNToOne(n - 1)
}

//Print Linearly from 1 to N (By Backtracking)
func printLinearlyFromOneToN(start int, n int) string {
	if start < 1 {
		return ""
	}
	fmt.Printf("Name %d\n", start)
	return printLinearlyFromOneToN(start - 1, n)

}

//Print from N to 1 (By Backtrack)
func printFromNToOneBacktrack(end int, n int) string {
	if end > n {
		return ""
	}
	fmt.Printf("Name %d\n", end)
	return printFromNToOneBacktrack(end + 1, n)
}

//Using functional recursion
func summationOfNumber(num int) int {
	if num == 0 {
		return 0
	}
	return num + summationOfNumber(num - 1)
}

func multOfNumber(n int) int {
	if n == 0 {
		return 1
	}
	return n * multOfNumber(n - 1)
}

func reverseArrUsingRecursionUsingTwoPointer(leftIndex int, rightIndex int, arr []int) []int {
	if leftIndex >= rightIndex {
		return arr
	}

	//Swap
	currentLeft := arr[leftIndex]
	currentRight := arr[rightIndex]
	arr[leftIndex] = currentRight
	arr[rightIndex] = currentLeft

	leftIndex++
	rightIndex--
	return reverseArrUsingRecursionUsingTwoPointer(leftIndex, rightIndex, arr)
}

func checkIfPalindromeUsingRecursion (startPoint int, str string) bool {
	endPoint := len(str) - startPoint - 1
	if str[startPoint] != str[endPoint] {
		return false
	}

	if len(str) % 2 == 0 && str[startPoint] == str[endPoint] && startPoint >= len(str)/2 {
		fmt.Println("Wow", startPoint, endPoint)
		return true
	}

	if len(str) % 2 != 0 && str[startPoint] == str[endPoint] {
		return true
	}
	return checkIfPalindromeUsingRecursion(startPoint + 1, str)
}

func fibonnaci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonnaci(n - 1) + fibonnaci(n - 2)
}