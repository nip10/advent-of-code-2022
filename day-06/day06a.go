package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("./input.txt")
	inputString := string(inputBytes)
	inputStringWithoutWindowsNewLines := strings.ReplaceAll(inputString, "\r\n", "\n")
	var index int = -1
	var dict = make(map[string]int)
	for charIndex := range inputStringWithoutWindowsNewLines[:len(inputStringWithoutWindowsNewLines)-4] {
		for i := 0; i < 4; i++ {
			dict[string(inputStringWithoutWindowsNewLines[charIndex+i])] += 1
		}
		if len(dict) == 4 {
			index = charIndex + 4
			break
		} else {
			dict = make(map[string]int)
		}
	}
	fmt.Println(index)
}

// Hardcoded solution
// if inputStringWithoutWindowsNewLines[i] != inputStringWithoutWindowsNewLines[i+1] &&
//  inputStringWithoutWindowsNewLines[i] != inputStringWithoutWindowsNewLines[i+2] &&
//  inputStringWithoutWindowsNewLines[i] != inputStringWithoutWindowsNewLines[i+3] &&
//  inputStringWithoutWindowsNewLines[i+1] != inputStringWithoutWindowsNewLines[i+2] &&
//   inputStringWithoutWindowsNewLines[i+1] != inputStringWithoutWindowsNewLines[i+3] &&
// 	inputStringWithoutWindowsNewLines[i+2] != inputStringWithoutWindowsNewLines[i+3] {
// 	index = i + 4
// 	break
// }
