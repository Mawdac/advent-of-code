package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file := "input.txt"
	fmt.Printf("Parsing %s\n", file)

	biggestChonker := parseInput(file)

	fmt.Printf("Aaand the biggest chonker is %d calories!\n", biggestChonker)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(file string) int {
	// open the file, check for errors, remember to close
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	// get a reader for the file
	scanner := bufio.NewScanner(f)

	var topChonker, currentChonker int

	// for each line
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			currentChonker = 0
			fmt.Println("A new chonker appears...")
		default:
			chonk, err := strconv.Atoi(scanner.Text())
			check(err)
			currentChonker += chonk
			if currentChonker > topChonker {
				topChonker = currentChonker
				fmt.Printf("~ New top chonker! %d\n", topChonker)
			} else {
				fmt.Printf("Current chonker: %d Top chonker: %d\n", currentChonker, topChonker)
			}
		}
	}

	return topChonker
}
