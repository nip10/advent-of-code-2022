package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
we can use a LIFO stack

For this example
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3
 The resulting stacks should be
 [N] [Z]
 [D] [C] [M]
 [P]
 so the first index is the top of the stack
*/

// LIFO stack
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true
}

// function to reverse given arr without generating new one
// with the help of left and right pointers
func reverse(start, end int, arr []string) {
	for start < end {
		// using left and right pointers
		arr[start], arr[end] = arr[end], arr[start]
		// moving forward
		start++
		// moving backwards
		end--
	}
}

func main() {
	inputBytes, _ := os.ReadFile("./input.txt")
	inputString := string(inputBytes)
	inputStringWithoutWindowsNewLines := strings.ReplaceAll(inputString, "\r\n", "\n")
	blocks := strings.Split(inputStringWithoutWindowsNewLines, "\n\n")
	stacksLines := strings.Split(blocks[0], "\n")
	stackNumbersArr := strings.Split(stacksLines[len(stacksLines)-1], " ")
	numStacks, _ := strconv.Atoi(stackNumbersArr[len(stackNumbersArr)-1])
	var stacks []Stack = make([]Stack, numStacks)
	stacksLinesClean := stacksLines[:len(stacksLines)-1]
	for _, line := range stacksLinesClean {
		for i := 1; i < len(line); i += 4 {
			if string(line[i]) != " " {
				stacks[i/4].Push(string(line[i]))
			}
		}
	}
	// reverse order of crates in stacks
	// this could be done in the loop above
	for _, stack := range stacks {
		reverse(0, len(stack)-1, stack)
	}
	instructionLines := blocks[1]
	for _, line := range strings.Split(instructionLines, "\n") {
		// eg: "move 1 from 2 to 1"
		lineData := strings.Split(line, " ")
		numCratesToMove, _ := strconv.Atoi(lineData[1])
		from, _ := strconv.Atoi(lineData[3])
		to, _ := strconv.Atoi(lineData[5])
		var tempStack Stack
		for i := 0; i < numCratesToMove; i++ {
			crate, _ := stacks[from-1].Pop()
			tempStack.Push(crate)
		}
		for i := 0; i < numCratesToMove; i++ {
			crate, _ := tempStack.Pop()
			stacks[to-1].Push(crate)
		}
	}
	var res string
	for _, stack := range stacks {
		res += stack[len(stack)-1]
	}
	fmt.Println("Result", res)
}
