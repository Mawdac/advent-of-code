package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)

	elvesTotals := calorieTotals(dat)

	topElf := elvesTotals[0]
	topThreeCombined := elvesTotals[0] + elvesTotals[1] + elvesTotals[2]

	fmt.Printf("Top Elf: %d\nTop Three Combined: %d\n", topElf, topThreeCombined)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calorieTotals(elves []byte) []int {
	elfStrings := strings.Split(string(elves), "\n\n")
	allElfCalories := make([]int, len(elfStrings))

	for i, elf := range elfStrings {
		foods := strings.Split(strings.TrimSpace(elf), "\n")
		fmt.Printf("New Elf %d: \n", i)
		for j, food := range foods {
			calories, err := strconv.Atoi(food)
			check(err)
			allElfCalories[i] += calories
			fmt.Printf("Food item %d: %d calories \n", j+1, calories)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allElfCalories)))

	return allElfCalories
}
