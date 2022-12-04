package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputBytes, err := os.ReadFile("./input.txt")
	check(err)
	inputString := string(inputBytes)
	inputStringWithoutWindowsNewLines := strings.ReplaceAll(inputString, "\r\n", "\n")
	pairs := strings.Split(inputStringWithoutWindowsNewLines, "\n")

	result := 0
	for _, pair := range pairs {
		pairParts := strings.Split(pair, ",")
		sections1 := strings.Split(pairParts[0], "-")
		min1, _ := strconv.Atoi(sections1[0])
		max1, _ := strconv.Atoi(sections1[1])
		sections2 := strings.Split(pairParts[1], "-")
		min2, _ := strconv.Atoi(sections2[0])
		max2, _ := strconv.Atoi(sections2[1])
		if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
			result++
		}
	}

	fmt.Println(result)
}
