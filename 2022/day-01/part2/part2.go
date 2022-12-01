package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file := "../input.txt"
	fmt.Printf("Parsing %s\n", file)

	top3Chonkers := parseInput(file)

	fmt.Printf("Top 3 Chonkers:\n %d\n %d\n %d\n", top3Chonkers[0], top3Chonkers[1], top3Chonkers[2])
	fmt.Printf("~ Total between all 3 ~\n%d\n", top3Chonkers[0]+top3Chonkers[1]+top3Chonkers[2])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(file string) []int {
	// open the file, check for errors, remember to close
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	// get a reader for the file
	scanner := bufio.NewScanner(f)

	var allChonkers []int
	var currentChonker int

	// for each line
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			if currentChonker != 0 {
				allChonkers = append(allChonkers, currentChonker)
			}
			currentChonker = 0
		default:
			chonk, err := strconv.Atoi(scanner.Text())
			check(err)
			currentChonker += chonk
		}
	}

	// append the last chonker left... TODO: do this better
	allChonkers = append(allChonkers, currentChonker)

	sort.Sort(sort.Reverse(sort.IntSlice(allChonkers)))

	return allChonkers
}
