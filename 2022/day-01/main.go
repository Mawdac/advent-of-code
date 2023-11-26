package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	elves := parseElvesFromFile(file)
	packs := parseCalories(elves)

	fmt.Printf("Processed %v elves!\n", len(packs))
	if len(os.Args) > 1 {
		for i, v := range packs {
			fmt.Print("Elf: ", i, " ")
			fmt.Print("Calories: ", v, "\n")
		}
		if os.Args[1] == "part1" || os.Args[1] == "all" {
			fmt.Printf("The elf with the most calories has %v\n", packs[0])
		}
		if os.Args[1] == "part2" || os.Args[1] == "all" {
			top3 := part2(packs)
			fmt.Printf("The top 3 elves have %v, %v and %v calories respectively\n", top3[0], top3[1], top3[2])
		}
	} else {
		fmt.Println("Try passing part1, part2, to the 'go run main.go' command to see solutions.")
	}
}

// Returns the elf with the most calories
func part1(packs []int) int {
	sorted := sortPacks(packs)
	return sorted[0]
}

// Returns the top 3 elves
func part2(packs []int) []int {
	sorted := sortPacks(packs)
	return []int{sorted[0], sorted[1], sorted[2]}
}

// Sort a []int into descending order and return a copy
func sortPacks(packs []int) []int {
	packsCopy := make([]int, len(packs))
	copy(packsCopy, packs)
	sort.Sort(sort.Reverse(sort.IntSlice(packsCopy)))
	return packsCopy
}

// Read each elfs pack as a string, splitting by 2 newlines
func parseElvesFromFile(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read %v: %v", file, err)
	}
	return strings.Split(string(f), "\n\n")
}

// Parse the raw elf packs string into a slice of the total calories per elf
func parseCalories(elves []string) []int {
	allElfPacks := make([]int, len(elves))

	for i, pack := range elves {
		snacks := strings.Fields(pack)
		if len(snacks) == 0 {
			log.Println("Empty snack pack:", snacks)
			continue
		}
		for _, snack := range snacks {
			calories, err := strconv.Atoi(snack)
			if err != nil {
				log.Printf("Could not convert %v to int: %v", snack, err)
				continue
			}
			allElfPacks[i] += calories
		}
	}
	return allElfPacks
}
