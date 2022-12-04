// https://adventofcode.com/2022/day/1
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Match struct {
	Opponent string
	Self     string
}

func loadMatches(path string) []Match {
	csvfile, _ := os.Open(path)
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.Comma = ' '

	rawCSVdata, _ := reader.ReadAll()

	var matches []Match
	for _, record := range rawCSVdata {
		match := Match{Opponent: record[0], Self: record[1]}
		matches = append(matches, match)
	}
	return matches
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	pointValues := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	matchupValues := map[string]string{"A": "Y", "B": "Z", "C": "X", "X": "B", "Y": "C", "Z": "A"}
	matches := loadMatches("input.csv")
	var maxScore []int
	var totalScore []int

	for _, match := range matches {
		switch pointValues[match.Self] - pointValues[match.Opponent] {
		case 1, -2: // Win
			maxScore = append(maxScore, pointValues[match.Self]+6)
		case -1, 2: // Lose
			maxScore = append(maxScore, pointValues[match.Self])
		case 0: // Draw
			maxScore = append(maxScore, pointValues[match.Self]+3)
		}
		switch match.Self {
		case "Z": // Win. Get winning matchup value.
			totalScore = append(totalScore, pointValues[matchupValues[match.Opponent]]+6)
		case "X": // Lose
			totalScore = append(totalScore, pointValues[matchupValues[matchupValues[match.Opponent]]]+0)
		case "Y": // Draw
			totalScore = append(totalScore, pointValues[match.Opponent]+3)
		}
	}

	fmt.Printf("When it was all over and the dust had settled, you final score was %v...\nIf you had cheated correctly, you would have secured %v points.\n", sum(maxScore), sum(totalScore))
}
