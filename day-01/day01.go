package main

import (
	"fmt"
	"os"
	"sort"
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
	elfsInventories := strings.Split(inputStringWithoutWindowsNewLines, "\n\n")

	var elfsCalories []int

	for _, elfInventory := range elfsInventories {
		elfCaloriesSum := 0
		elfCalories := strings.Split(elfInventory, "\n")
		for _, ingredientCalories := range elfCalories {
			caloriesInt, err := strconv.Atoi(ingredientCalories)
			check(err)
			elfCaloriesSum += caloriesInt
		}
		elfsCalories = append(elfsCalories, elfCaloriesSum)
	}

	sort.Ints(elfsCalories)
	highestThreeCalorieElves := elfsCalories[len(elfsCalories)-3:]

	fmt.Println("Part A", elfsCalories[len(elfsCalories)-1])
	fmt.Println("Part B", highestThreeCalorieElves[0]+highestThreeCalorieElves[1]+highestThreeCalorieElves[2])
}
