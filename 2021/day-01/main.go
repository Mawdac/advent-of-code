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

func measureSweeps(sweeps []int) int {
	var increases int

	for i, v := range sweeps {
		if i == 0 {
			fmt.Printf("Skipping first iteration %d\n", v)
			continue
		} else if v > sweeps[i-1] {
			increases++
		}
		fmt.Printf("Increases: %d - Previous: %d - Current %d\n", increases, sweeps[i-1], v)
	}

	return increases
}

func readInput(file string) []int {
	var values []int

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, x)
	}

	return values
}
