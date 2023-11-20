package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rucksack := readRucksack("input.txt")
	var part1Total int
	var part2Total int

	fmt.Println("-=-=- Part 1 -=-=-")
	for i := 0; i < len(rucksack)-1; i++ {
		sack1, sack2 := halfString(rucksack[i])
		common := compareSacks(sack1, sack2)
		priority := getPriority([]byte(common)[0])
		part1Total += int(priority)
	}
	fmt.Println("Solution:", part1Total)

	fmt.Println("-=-=- Part 2 -=-=-")
	// for each group of 3 elves
	for i := 0; i < len(rucksack)-1; i += 3 {
		elf1, elf2, elf3 := rucksack[i], rucksack[i+1], rucksack[i+2]
		fmt.Println("Elf 1:", elf1)
		fmt.Println("Elf 2:", elf2)
		fmt.Println("Elf 3:", elf3)
		badge := findBadge(elf1, elf2, elf3)
		fmt.Println("Badge:", badge)
		priority := getPriority([]byte(badge)[0])
		part2Total += int(priority)
		fmt.Println(i, priority)
	}

	fmt.Println("Solution 2:", part2Total)
}

func readRucksack(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	return strings.Split(string(data), "\n")
}

func getPriority(b byte) int64 {
	val, err := strconv.ParseInt(fmt.Sprintf("%x", b), 16, 32)
	if err != nil {
		fmt.Println("Error converting:", err)
		os.Exit(1)
	}
	// return val
	if 65 <= val && val <= 90 {
		return val - 38
	} else if 97 <= val && val <= 122 {
		return val - 96
	} else {
		return 0
	}
}

func halfString(s string) (string, string) {
	b := []byte(s)
	s1 := string(b[:len(b)/2])
	s2 := string(b[len(b)/2:])
	return s1, s2
}

func findBadge(s1 string, s2 string, s3 string) string {
	s1Slice := []byte(s1)
	s2Slice := []byte(s2)
	s3Slice := []byte(s3)

	for _, v1 := range s1Slice {
		for _, v2 := range s2Slice {
			if v1 == v2 {
				for _, v3 := range s3Slice {
					if v2 == v3 {
						return string(v1)
					}
				}
			}
		}
	}
	// TODO: Return an error instead of just blank string
	return ""
}

func compareSacks(s1 string, s2 string) string {
	s1Slice := []byte(s1)
	s2Slice := []byte(s2)

	for _, v1 := range s1Slice {
		for _, v2 := range s2Slice {
			if v1 == v2 {
				return string(v1)
			}
		}
	}
	// TODO: Return an error instead of just blank string
	return ""
}
