package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var lines []string
	if len(os.Args) > 1 {
		lines = readLines(os.Args[1])
		fmt.Printf("Read %v lines from file\n", len(lines))
	}

	if len(os.Args) > 2 {
		for i, v := range lines {
			fmt.Println("Line: ", i+1)
			fmt.Println("Value:", v)
		}
		if os.Args[2] == "part1" || os.Args[2] == "all" {
			total := part1(lines)
			fmt.Println("~ Part 1 Solution ~")
			fmt.Printf("The sum of all calibration values is %v\n", total)
		}
		if os.Args[2] == "part2" || os.Args[2] == "all" {
			total2 := part2(lines)
			fmt.Println("\n~ Part 2 Solution ~")
			fmt.Printf("The sum of all calibration values is %v\n", total2)
		}
	} else {
		fmt.Println("Missing command flags")
		fmt.Println("Try this: 'go run main.go input.txt all/part1/part2'")
	}
}

func readLines(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read %v: %v", file, err)
	}

	lines := strings.TrimSuffix(string(f), "\n")
	return strings.Split(lines, "\n")
}

func part1(lines []string) int {
	total := 0
	for _, v := range lines {
		total += parseOnlyNumbersFromCalibration(v)
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, v := range lines {
		total += parseNumbersAndStringsFromCalibration(v)
	}
	return total
}

func parseOnlyNumbersFromCalibration(line string) int {
	runes := make([]rune, 0)
	firstAndLast := make([]rune, 2)
	for _, v := range line {
		if unicode.IsNumber(v) {
			runes = append(runes, v)
		}
	}

	// if there's nothing found, return
	if len(runes) == 0 {
		fmt.Printf("No integers found in %v\n", line)
		return 0
	}
	firstAndLast[0] = runes[0]

	if len(runes) > 1 {
		firstAndLast[1] = runes[len(runes)-1]
	} else {
		firstAndLast[1] = runes[0]
	}
	fmt.Printf("runes: %v\nfirst and last:%v\n", runes, firstAndLast)

	calibration, err := strconv.Atoi(string(firstAndLast))
	if err != nil {
		log.Printf("Could not convert %v to int: %v\n", string(runes), err)
	}

	return calibration
}

var numberWords = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var hotRunes = []rune{'o', 't', 'f', 's', 'e', 'n'}

func isRuneHot(r rune) bool {
	for _, v := range hotRunes {
		if r == v {
			return true
		}
	}
	return false
}

func parseNumbersAndStringsFromCalibration(line string) int {
	parsedString := ""
	splitNumbersAndStrings := splitByNumbersAndStrings(line)

	for _, v := range splitNumbersAndStrings {
		// if it's a number append it to our returnString
		if !unicode.IsLetter(rune(v[0])) {
			parsedString += string(v)
		} else {
			// if it's not worth checking...
			if len(v) < 3 {
				continue // don't bother with it
			}

			parsedNumbers := parseNumbersFromStrings(v)
			for _, v := range parsedNumbers {
				parsedString += v
			}
		}
	}

	firstAndLast := string(parsedString[0]) + string(parsedString[len(parsedString)-1])

	calibration, err := strconv.Atoi(firstAndLast)
	if err != nil {
		log.Printf("Could not convert %v to int: %v\n", string(firstAndLast), err)
	}
	return calibration
}

func splitByNumbersAndStrings(line string) []string {
	sliceLine := make([]string, 0, len(line))
	sliceIndex := 0
	lastIsLetter := false
	// for each rune
	for _, v := range line {
		// if it's a number
		if unicode.IsNumber(v) {
			if lastIsLetter {
				sliceIndex++
			}
			sliceLine = append(sliceLine, string(v))
			sliceIndex++
		} else {
			// this is the first letter of a string
			if !lastIsLetter {
				sliceLine = append(sliceLine, string(v))
			} else {
				sliceLine[sliceIndex] += string(v)
			}
		}
		lastIsLetter = unicode.IsLetter(v)
	}
	return sliceLine
}

func parseNumbersFromStrings(text string) []string {
	value := make([]string, 0)

	for i, v := range text {
		if !isRuneHot(v) {
			continue // if it doesn't start with a letter we care about, move on
		}
		if len(text[i:]) >= 3 && numberWords[text[i:i+3]] != "" {
			parsedNumber := numberWords[text[i:i+3]]
			value = append(value, parsedNumber)
		}
		if len(text[i:]) >= 4 && numberWords[text[i:i+4]] != "" {
			parsedNumber := numberWords[text[i:i+4]]
			value = append(value, parsedNumber)
		}
		if len(text[i:]) >= 5 && numberWords[text[i:i+5]] != "" {
			parsedNumber := numberWords[text[i:i+5]]
			value = append(value, parsedNumber)
		}
	}

	return value
}
