package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	fmt.Print("Parsing input for total score...")
	totalA, totalB := totalScore(data)
	fmt.Printf("Part 1:%d\nPart 2:%d\n", totalA, totalB)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func totalScore(input []byte) (int, int) {
	rounds := strings.Split(string(input), "\n")
	var partOneTotal, partTwoTotal int

	for _, round := range rounds {
		choices := strings.Split(strings.TrimSpace(round), " ")
		if len(choices) < 2 {
			break
		}
		fmt.Printf("%v vs %v\n", choices[0], choices[1])

		// TODO: find a better way...
		switch choices[0] {
		case "A":
			switch choices[1] {
			case "X":
				partOneTotal += 4 // rock=1, draw=3
				partTwoTotal += 3 // lose=0, scissors=3
			case "Y":
				partOneTotal += 8 // paper=2, win=6
				partTwoTotal += 4 // draw=3, rock=1
			case "Z":
				partOneTotal += 3 // scissors=3, lose=0
				partTwoTotal += 8 // win=6, paper=2
			}
		case "B":
			switch choices[1] {
			case "X":
				partOneTotal += 1 // rock=1, lose=0
				partTwoTotal += 1 // lose=0, rock=1
			case "Y":
				partOneTotal += 5 // paper=2, draw=3
				partTwoTotal += 5 // draw=3, paper=2
			case "Z":
				partOneTotal += 9 // scissors=3, win=6
				partTwoTotal += 9 // win=6, scissors=3
			}
		case "C":
			switch choices[1] {
			case "X":
				partOneTotal += 7 // rock=1, win=6
				partTwoTotal += 2 // lose=0, paper=2
			case "Y":
				partOneTotal += 2 // paper=2, lose=0
				partTwoTotal += 6 // draw=3, scissors=3
			case "Z":
				partOneTotal += 6 // scissors=3, draw=3
				partTwoTotal += 7 // win=6, rock=1
			}
		}
	}

	return partOneTotal, partTwoTotal
}
