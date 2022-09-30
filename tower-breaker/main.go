package main

import (
	"fmt"
)

func TowerBreaker(tower, tall int32) {
	win := map[int32]int32{
		1: 1,
	}

	originalTall := tall

	play := true
	player := 1
	var breaker int32 = 1

	for play {
		if breaker < 1 || breaker >= tall {
			if player%2 != 0 {
				win[1] = 1
				play = false
			} else {
				win[1] = 2
				play = false
			}
		}

		tall -= breaker
		player++
		breaker++
	}

	players := win[1]

	fmt.Println(FinalResult(int(originalTall), int(players)))

}

func FinalResult(originalTall int, players int) string {
	if originalTall%2 == 0 {
		return Message(players)
	} else {
		return Message(players)
	}
}

func Message(players int) string {
	if players == 1 {
		return "Player 1 WIN"
	} else {
		return "Player 1 LOSE"
	}
}

func main() {
	TowerBreaker(2, 6)
}
