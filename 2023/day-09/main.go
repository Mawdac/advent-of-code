package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, 2023 day 9!")

	sensorValues := readSensorValues("test.txt")

	for i, v := range sensorValues {
		fmt.Printf("Line %v: %v\n", i, v)
	}

	fmt.Printf("The solution to part 1 is %v\n", part1(sensorValues))
}

func part1(values [][]int) int {
	sumOfValues := 0

	for i, v := range values {
		fmt.Printf("Find the next value for %v...\n", v)
		newValue := findNextDifference(v) + v[len(v)-1]
		sumOfValues += newValue
		fmt.Printf("Next value for %v is %v\n", i, newValue)
	}

	return sumOfValues
}

func findNextDifference(values []int) int {
	differences := make([]int, len(values)-1)
	for i := 0; i < len(values)-1; i++ {
		differences[i] = absDiffInt(values[i], values[i+1])
	}

	fmt.Printf("Differences for %v: %v\n", values, differences)

	if checkForZeros(differences) {
		return differences[0]
	} else {
		return findNextDifference(differences) + differences[len(differences)-1]
	}
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func checkForZeros(values []int) bool {
	for i := 0; i < len(values); i++ {
		if values[i] != 0 {
			return false
		}
	}
	return true
}

func readSensorValues(file string) [][]int {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v\n", file, err)
	}

	lines := strings.Split(trim(string(f)), "\n")
	returnSensorValues := make([][]int, len(lines))

	for i, l := range lines {
		returnSensorValues[i] = parseValues(l)
	}
	return returnSensorValues
}

func parseValues(line string) []int {
	returnSlice := []int{}

	numberStrings := strings.Split(line, " ")

	for _, n := range numberStrings {
		value, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("Couldn't convert %v, %v\n", n, err)
		}
		returnSlice = append(returnSlice, value)
	}

	return returnSlice
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
