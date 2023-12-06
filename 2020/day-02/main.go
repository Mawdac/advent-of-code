package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	minMax   []int
	password []rune
	letter   rune
}

func main() {
	fmt.Println("Hello, world!")

	entries := readEntries("input.txt")
	for _, v := range entries {
		fmt.Printf("min:%v, max:%v, letter:%v, password:%v\n", v.minMax[0], v.minMax[1], string(v.letter), string(v.password))
	}

	fmt.Printf("The solution to part 1 is %v\n", part1(entries))
	fmt.Printf("The solution to part 2 is %v\n", part2(entries))
}

func part1(e []entry) int {
	returnTotal := 0
	for _, v := range e {
		valid := checkPolicy(v)
		if valid {
			returnTotal++
		}
	}
	return returnTotal
}

func part2(e []entry) int {
	returnTotal := 0
	for _, v := range e {
		valid := checkNewPolicy(v)
		if valid {
			returnTotal++
		}
	}
	return returnTotal
}

func checkNewPolicy(e entry) bool {
	var firstLetter, secondLetter bool

	if e.password[e.minMax[0]-1] == e.letter {
		firstLetter = true
	}
	if e.password[e.minMax[1]-1] == e.letter {
		secondLetter = true
	}

	if (!firstLetter && !secondLetter) || (firstLetter && secondLetter) {
		return false
	}
	return true
}

func checkPolicy(e entry) bool {
	runeCount := 0
	for _, r := range e.password {
		if r != e.letter {
			continue
		}
		runeCount++
	}

	if runeCount < e.minMax[0] || runeCount > e.minMax[1] {
		return false
	}
	return true
}

func readEntries(file string) []entry {
	returnEntries := make([]entry, 0)
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read the file: %v", err)
	}

	lines := strings.Split(trim(string(f)), "\n")

	for _, v := range lines {
		returnEntries = append(returnEntries, parseLine(v))
	}
	return returnEntries
}

func parseLine(line string) entry {
	returnEntry := entry{}

	lineSlice := strings.Split(trim(line), " ")
	minAndMax := strings.Split(lineSlice[0], "-")
	returnEntry.letter = []rune(lineSlice[1])[0]
	returnEntry.password = []rune(lineSlice[2])

	min, err := strconv.Atoi(minAndMax[0])
	if err != nil {
		log.Fatalf("Couldn't read the file: %v\n", err)
	}
	max, err := strconv.Atoi(minAndMax[1])
	if err != nil {
		log.Fatalf("Couldn't read the file: %v\n", err)
	}

	returnEntry.minMax = append(returnEntry.minMax, min, max)

	return returnEntry
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
