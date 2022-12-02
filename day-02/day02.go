package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// A/X => Rock
// B/Y => Paper
// C/Z => Scissors

/*
Other ways to solve this, or tweaks:
- use maps for individual scores
- use map with hardcoded scores (A/X = 3+1, A/Y = 6+2, A/Z = 0+3)
- use switch case
*/

func main() {
	inputBytes, err := os.ReadFile("./input.txt")
	check(err)
	inputString := string(inputBytes)
	inputStringWithoutWindowsNewLines := strings.ReplaceAll(inputString, "\r\n", "\n")
	rounds := strings.Split(inputStringWithoutWindowsNewLines, "\n")

	var points int = 0
	for _, round := range rounds {
		playersChoices := strings.Split(round, " ")
		choice1 := playersChoices[0]
		myChoice := playersChoices[1]
		// points by result
		if choice1 == "A" && myChoice == "X" || choice1 == "B" && myChoice == "Y" || choice1 == "C" && myChoice == "Z" {
			// draw
			points += 3
		} else if choice1 == "A" && myChoice == "Y" || choice1 == "B" && myChoice == "Z" || choice1 == "C" && myChoice == "X" {
			// win
			points += 6
		} else {
			// lose
			points += 0
		}
		// points by choice
		if myChoice == "X" {
			points += 1
		} else if myChoice == "Y" {
			points += 2
		} else if myChoice == "Z" {
			points += 3
		}
	}
	fmt.Println("Result:", points)

	var pointsB int = 0
	for _, round := range rounds {
		playersChoices := strings.Split(round, " ")
		choice1 := playersChoices[0]
		gameResult := playersChoices[1]
		// points by result
		if gameResult == "X" {
			// lose
			pointsB += 0
		} else if gameResult == "Y" {
			// draw
			pointsB += 3
		} else if gameResult == "Z" {
			// win
			pointsB += 6
		}
		// points by choice
		if choice1 == "A" && gameResult == "Y" || choice1 == "B" && gameResult == "X" || choice1 == "C" && gameResult == "Z" {
			// I play rock
			pointsB += 1
		} else if choice1 == "A" && gameResult == "Z" || choice1 == "B" && gameResult == "Y" || choice1 == "C" && gameResult == "X" {
			// I play paper
			pointsB += 2
		} else {
			// I play scissors
			pointsB += 3
		}
	}
	fmt.Println("Result B:", pointsB)
}
