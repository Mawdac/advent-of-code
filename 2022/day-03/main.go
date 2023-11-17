package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rucksack := readRucksack("input.txt")
	var total int
	for i := 0; i < len(rucksack)-1; i++ {
		sack1, sack2 := halfString(rucksack[i])
		common := compareSacks(sack1, sack2)
		priority := getPriority([]byte(common)[0])
		total += int(priority)
		fmt.Println(i, priority)
	}
	fmt.Println("Solution:", total)
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
