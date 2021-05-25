package scrabble

import "strings"

/*
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/

type ScoreSchema map[string]int

func Score(word string) int {
	if len(word) <= 0 {
		return 0
	}
	totalScore := 0
	wordLowerCase := strings.ToLower(word)
	var scoreMap = ScoreSchema{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvmy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}
	for scoreKey, scoreValue := range scoreMap {
		totalScore += pointsCalculator(wordLowerCase, scoreKey, scoreValue)
	}
	return totalScore
}

func pointsCalculator(word string, scoreBoard string, scoreValue int) int {
	totalScore := 0
	for i := 0; i < len(word); i++ {
		for z := 0; z < len(scoreBoard); z++ {
			if word[i] == scoreBoard[z] {
				totalScore += scoreValue
			}
		}
	}
	return totalScore
}
