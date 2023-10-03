package main

import (
	"fmt"
	"os"
	"strings"
)

func getOutcomeSignScore(outcome int, sign string) int {

	a := []string{"A", "B", "C"}

	score := 0

	for i, v := range a {
		if sign == v {
			if outcome == 3 {
				score = i + 1
			}
			if outcome == 6 {
				if i+1 == len(a) {
					score = 1
				} else {
					score = i + 2
				}
			}

			if outcome == 0 {
				if i-1 == -1 {
					score = 3
				} else {
					score = i
				}
			}
		}
	}

	return score
}

func main() {
	path := os.Args[1]
	outcomeScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	file, _ := os.ReadFile(path)

	s := string(file)

	sArr := strings.Split(s, "\n")

	tournamentScore := 0

	for i := 0; i < len(sArr)-1; i++ {
		fmt.Println(sArr[i])
		opponentSign := string(sArr[i][0])
		outcome := string(sArr[i][2])

		signScore := getOutcomeSignScore(outcomeScoreMap[outcome], opponentSign)
		outcomeScore := outcomeScoreMap[outcome]

		tournamentScore = tournamentScore + outcomeScore + signScore

	}

	fmt.Println(tournamentScore)
}
