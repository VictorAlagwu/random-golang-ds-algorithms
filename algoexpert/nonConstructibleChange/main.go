package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{1,1,1,1,1}
	fmt.Println("Non Constructable Change:", nonConstructibleChange(coins))

}

func nonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	amountOfChange := 0

	for _, value := range coins {
		if value > amountOfChange {
			amountOfChange += 1
		}

		amountOfChange += value
	}
	return amountOfChange
}