package main

const HomeTeamWon = 1

func main(){

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

