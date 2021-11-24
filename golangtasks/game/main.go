package main

import (
	"fmt"
	"os"
)

func main() {

	var n, m int
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	fmt.Println(findWinner(n, m))

}

func findWinner(n, m int) int {
	if m == 0 {
		return n - 1
	} else if n == 0 {
		fmt.Println("Cannot play the game without players!")
		os.Exit(0)
	}

	players := make([]int, n)
	outOfGamePlayer := 0
	counter := 0

	for i := 0; i < len(players); i++ {
		if players[i] == 0 {
			counter++
			if counter == m {
				players[i] = 1
				outOfGamePlayer++
				counter = 0
				if i == len(players)-1 {
					i = -1
				} else {
					continue
				}
			}
		}
		if i == len(players)-1 {
			i = -1
		}
		if outOfGamePlayer == len(players)-1 {
			break
		}
	}
	return getIndexOfLastPlayer(players)
}

func getIndexOfLastPlayer(players []int) int {
	var indexOfLastPlayer int
	for i := 0; i < len(players); i++ {
		if players[i] == 0 {
			indexOfLastPlayer = i
		}
	}
	return indexOfLastPlayer + 1
}
