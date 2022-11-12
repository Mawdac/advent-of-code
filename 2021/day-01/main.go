package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput("input.txt")

	fmt.Printf("Total increases on sweep depth: %d\n", measureSweeps(input))
}

// count the amount of increases between iterations for a given []int
func measureSweeps(sweeps []int) int {
	var positiveCount int
	var currRange, prevRange []int

	// for each sweep
	for i := range sweeps {
		// if it's the first time, skip
		if i == 0 {
			currRange = sweeps[i : i+3]
			continue
		}

		// if there's not enough elements left, we're done
		if i+3 > len(sweeps) {
			break
		}

		// set the previous range, then update the current one
		prevRange = currRange
		currRange = sweeps[i : i+3]

		// if the sum of the current is more than the sum of the previous
		if sum(currRange) > sum(prevRange) {
			positiveCount++
		}
	}

	return positiveCount
}

// calculate the sum of a slice of integers
func sum(sweeps []int) int {
	total := 0
	for _, v := range sweeps {
		total += v
	}
	return total
}

// reads the input from a file and converts it to an array of integers
func readInput(file string) []int {
	var values []int

	// open the file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close it
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// for each line
	for scanner.Scan() {
		// conver it to an int
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// append it to the return array
		values = append(values, x)
	}

	return values
}
