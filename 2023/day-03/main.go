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
	schematic := readEngineSchematic("input.txt")
	for _, v := range schematic {
		fmt.Println(string(v))
	}
	fmt.Println("The solution to part 1 is", part1(schematic))
	fmt.Println("The solution to part 2 is", part2(schematic))
}

func part1(schematic [][]rune) int {
	returnTotal := 0
	for row, r := range schematic {
		for column := 0; column < len(r)-1; column++ {
			if unicode.IsNumber(schematic[row][column]) {
				numLength := checkNumberLength(row, column, schematic)
				if isPartNumber(row, column, numLength, schematic) {
					partNumber, err := strconv.Atoi(string(schematic[row][column : column+numLength]))
					if err != nil {
						log.Fatalf("Couldn't convert %v: %v", string(schematic[row][column:numLength]), err)
					}
					returnTotal += partNumber
				}
				column += numLength
			}
		}
	}
	return returnTotal
}

func part2(schematic [][]rune) int {
	returnTotal := 0
	for row, r := range schematic {
		for column := 0; column < len(r)-1; column++ {
			// check if it's a gear
			if schematic[row][column] == 42 {
				// check if it's surrounded by 2 numbers
				returnTotal += validateGear(row, column, schematic)
			}
		}
	}
	return returnTotal
}

func validateGear(row int, column int, schematic [][]rune) int {
	leftSpace, rightSpace := 0, 0
	parts := []int{}

	fmt.Println("Validating gear...")

	if column != 0 {
		leftSpace = 1
		if unicode.IsNumber(schematic[row][column-1]) {
			fmt.Println("Checking left side...", string(schematic[row][column-1]))
			newPart := discoverPartNumber(row, column-1, schematic)
			if newPart != 0 {
				parts = append(parts, newPart)
			}
		}
	}
	if column+1 < len(schematic[row]) {
		rightSpace = 1
		if unicode.IsNumber(schematic[row][column+1]) {
			fmt.Println("Checking right side...", string(schematic[row][column+1]))
			newPart := discoverPartNumber(row, column+1, schematic)
			if newPart != 0 {
				parts = append(parts, newPart)
			}
		}
	}

	if row != 0 {
		for i, v := range schematic[row-1][column-leftSpace : column+rightSpace+1] {
			if unicode.IsNumber(v) {
				fmt.Println("Checking above...", string(schematic[row-1][column-leftSpace+i]))
				// check for left/right space
				if leftSpace == 1 && rightSpace == 1 {
					// check for 2 part numbers in the row
					isTwofer := unicode.IsNumber(schematic[row-1][column-1]) && !unicode.IsNumber(schematic[row-1][column]) && unicode.IsNumber(schematic[row-1][column+1])
					if isTwofer {
						fmt.Println("It's a towfer!")
						p1 := discoverPartNumber(row-1, column-1, schematic)
						p2 := discoverPartNumber(row-1, column+1, schematic)
						parts = append(parts, p1, p2)
						break
					} else {
						newPart := discoverPartNumber(row-1, column-leftSpace+i, schematic)
						if newPart != 0 {
							parts = append(parts, newPart)
							break
						}
					}
				}
			}
		}
	}
	if row+1 < len(schematic) {
		for i, v := range schematic[row+1][column-leftSpace : column+rightSpace+1] {
			if unicode.IsNumber(v) {
				fmt.Println("Checking below...", string(schematic[row+1][column-leftSpace+i]))
				if leftSpace == 1 && rightSpace == 1 {
					// check for 2 part numbers in the row
					isTwofer := unicode.IsNumber(schematic[row+1][column-1]) && !unicode.IsNumber(schematic[row+1][column]) && unicode.IsNumber(schematic[row+1][column+1])
					if isTwofer {
						fmt.Println("It's a towfer!")
						p1 := discoverPartNumber(row+1, column-1, schematic)
						p2 := discoverPartNumber(row+1, column+1, schematic)
						parts = append(parts, p1, p2)
						break
					} else {
						newPart := discoverPartNumber(row+1, column-leftSpace+i, schematic)
						if newPart != 0 {
							parts = append(parts, newPart)
							break
						}
					}
				}
			}
		}
	}

	if len(parts) == 2 {
		fmt.Printf("Valid gear found! %v and %v surround %v:%v\n", parts[0], parts[1], row, column)
		return parts[0] * parts[1]
	}

	fmt.Printf("Not a valid gear: %v:%v\n", row, column)
	return 0
}

func discoverPartNumber(rowStart int, columnStart int, schematic [][]rune) int {
	numberRunes := make([]rune, 0)

	// left first
	for i := columnStart; i+1 > 0 && unicode.IsNumber(schematic[rowStart][i]); i-- {
		numberRunes = append([]rune{schematic[rowStart][i]}, numberRunes...)
	}
	// then right
	for i := columnStart + 1; i != len(schematic[rowStart]) && unicode.IsNumber(schematic[rowStart][i]); i++ {
		numberRunes = append(numberRunes, schematic[rowStart][i])
	}

	numberString := ""
	for _, v := range numberRunes {
		numberString += string(v)
	}
	returnNum, err := strconv.Atoi(string(numberString))
	if err != nil {
		log.Fatalf("Couldn't convert %v: %v", string(numberString), err)
	}

	fmt.Println("Discovered", returnNum)

	return returnNum
}

func isPartNumber(rowStart int, columnStart int, length int, schematic [][]rune) bool {
	surroundCharacters := make([]rune, 0)
	leftSpace, rightSpace := 0, 0

	if columnStart != 0 {
		leftSpace = 1
		surroundCharacters = append(surroundCharacters, schematic[rowStart][columnStart-1])
	}
	if columnStart+length+1 < len(schematic[rowStart]) {
		rightSpace = 1
		surroundCharacters = append(surroundCharacters, schematic[rowStart][columnStart+length])
	}

	if rowStart != 0 {
		surroundCharacters = append(surroundCharacters, schematic[rowStart-1][columnStart-leftSpace:columnStart+length+rightSpace]...)
	}
	if rowStart+1 < len(schematic) {
		surroundCharacters = append(surroundCharacters, schematic[rowStart+1][columnStart-leftSpace:columnStart+length+rightSpace]...)
	}

	for _, v := range surroundCharacters {
		if v != 46 && !unicode.IsNumber(v) {
			return true
		}
	}

	return false
}

func checkNumberLength(rowStart int, columnStart int, schematic [][]rune) int {
	returnLength := 0

	for _, v := range schematic[rowStart][columnStart:] {
		if !unicode.IsNumber(v) {
			return returnLength
		}
		returnLength += 1
	}
	return returnLength
}

func readEngineSchematic(file string) [][]rune {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read %v: %v", file, err)
	}

	trimmedFile := strings.TrimSuffix(string(f), "\n")
	lines := strings.Split(trimmedFile, "\n")
	returnSchematic := make([][]rune, len(lines))

	for i, v := range lines {
		returnSchematic[i] = []rune(strings.TrimSpace(v))
	}

	return returnSchematic
}
