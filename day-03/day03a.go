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

var alphaDict = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

// The Elf that did the packing failed to follow this rule for exactly one item type per rucksack

func main() {
	inputBytes, err := os.ReadFile("./input.txt")
	check(err)
	inputString := string(inputBytes)
	inputStringWithoutWindowsNewLines := strings.ReplaceAll(inputString, "\r\n", "\n")
	rucksacks := strings.Split(inputStringWithoutWindowsNewLines, "\n")

	var priorities int = 0
	for _, rucksack := range rucksacks {
		halfLength := len(rucksack) / 2
		var dict = make(map[string]int)
		for i := 0; i < halfLength; i++ {
			dict[rucksack[i:i+1]] = 1
		}
		for i := halfLength; i < len(rucksack); i++ {
			if dict[rucksack[i:i+1]] == 1 {
				priorities += alphaDict[rucksack[i:i+1]]
				break
			}
		}
	}
	fmt.Println("Result:", priorities)
}
