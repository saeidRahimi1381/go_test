//https://exercism.org/tracks/go/exercises/scrabble-score

package main

import "strings"

func main() {

}

func Score(word string) int {
	Score := 0
	for _, c := range strings.ToLower(word) {
		Score := ScoreOf(rune(c))
	}
}

func ScoreOf(c rune) int {
	switch c {

	}
}
