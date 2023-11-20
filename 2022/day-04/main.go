package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	data := readCSV(file)
	part1Total := 0
	part2Total := 0

	for _, v := range data {
		elf1Lower, elf1Upper := splitRange(v[0])
		elf2Lower, elf2Upper := splitRange(v[1])
		elf1 := []int{elf1Lower, elf1Upper}
		elf2 := []int{elf2Lower, elf2Upper}

		if containsRange(elf1, elf2) || containsRange(elf2, elf1) {
			part1Total += 1
		}

		if hasOverlap(elf1, elf2) {
			part2Total += 1
		}
	}
	fmt.Println("Part 1 Solution:", part1Total)
	fmt.Println("Part 2 Solution:", part2Total)
}

func readCSV(file string) [][]string {
	data, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	reader := csv.NewReader(data)

	records, _ := reader.ReadAll()
	return records
}

func splitRange(width string) (int, int) {
	splitString := strings.Split(width, "-")
	lower, _ := strconv.Atoi(splitString[0])
	upper, _ := strconv.Atoi(splitString[1])
	return lower, upper
}

func containsRange(range1 []int, range2 []int) bool {
	if range1[0] <= range2[0] {
		if range1[1] >= range2[1] {
			fmt.Println(range1, "contains", range2)
			return true
		}
	}
	return false
}

func hasOverlap(range1 []int, range2 []int) bool {
	if range1[0] > range2[0] {
		range1, range2 = range2, range1
	}
	return range1[1] >= range2[0]
}
