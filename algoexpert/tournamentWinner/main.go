package main

import (
	"fmt"
)

const HomeTeamWon = 1

func main(){
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0,0,1}
	fmt.Println("Tournament Winner: O(n)", tournamentWinner(competitions, results))
}
func tournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentWinner := ""
	arr := map[string]int{currentWinner: 0}

	for idx, value := range competitions {
		homeTeam, awayTeam := value[0], value[1]
		winningTeam := awayTeam
		if results[idx] == HomeTeamWon {
			winningTeam = homeTeam
		}
		arr[winningTeam] += 3
		if arr[winningTeam] > arr[currentWinner] {
			currentWinner = winningTeam
		}
	}
	return currentWinner
}

