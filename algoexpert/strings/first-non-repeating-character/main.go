package main

import "fmt"

func main() {
	word := "ababac"
	fmt.Println("First Non Repeating Character: ", firstNonRepeatingCharacter(word))
	fmt.Println("First Non Repeating Character: ", firstNonRepeatingCharacterII(word))
}

func firstNonRepeatingCharacter(str string) int {
	// Write your code here.
	//O(n^2) time complexity, O(1) space complexity
	currentIndex := -1
	hashTable := map[uint8]bool{}
	for i := 0; i < len(str); i++ {
		for j := 1; j < len(str); j++ {
			if str[i] == str[j] && i != j {
				hashTable[str[j]] = true
			}
		}
		if _, found := hashTable[str[i]]; found {
			continue
		}
		currentIndex = i
		if currentIndex != -1 {
			break
		}
	}
	return currentIndex
}

func firstNonRepeatingCharacterII(str string) int {
	// Write your code here.
	//O(n) time complexity, O(1) space complexity
	currentIndex := 0
	hashTable := map[uint8]int{}

	for currentIndex < len(str) {
		hashTable[str[currentIndex]] += 1
		currentIndex++
	}
	currentIndex = 0
	for currentIndex < len(str) {
		if hashTable[str[currentIndex]] == 1 {
			return currentIndex
		}
		currentIndex++
	}
	return -1
}