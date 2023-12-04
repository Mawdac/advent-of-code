package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type scratchcard struct {
	winners []int
	numbers []int
	count   int
}

func main() {
	var scratchcards []scratchcard
	if len(os.Args) > 1 {
		scratchcards = readScratchcards(os.Args[1])
		fmt.Printf("Read in %v scratchcards\n\n", len(scratchcards))
	}

	if len(os.Args) > 2 {
		for i, v := range scratchcards {
			fmt.Println("Game: ", i+1)
			fmt.Println("Sets:", v)
		}
		if os.Args[2] == "part1" || os.Args[2] == "all" {
			p1 := part1(scratchcards)
			fmt.Println("~ Part 1 Solution ~")
			fmt.Printf("The solution for part 1 is %v\n", p1)
		}
		if os.Args[2] == "part2" || os.Args[2] == "all" {
			p2 := part2(scratchcards)
			fmt.Println("~ Part 2 Solution ~")
			fmt.Printf("The solution for part 2 is %v\n", p2)
		}
	} else {
		fmt.Println("Missing command flags")
		fmt.Println("Try this: 'go run main.go input.txt all/part1/part2'")
	}
}

func part1(cards []scratchcard) int {
	total := 0

	for _, v := range cards {
		points := 0
		matches := checkMatches(v)
		for i := 0; i < matches; i++ {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
		total += points
	}

	return total
}

func part2(cards []scratchcard) int {
	total := 0
	// for each card
	for i, c := range cards {
		// check how many matches there are
		matches := checkMatches(c)
		// for each match, copy the next x(matches) cards
		for i2 := 0; i2 < matches; i2++ {
			cards[i+1+i2].count += c.count
		}
	}

	for _, v := range cards {
		total += v.count
	}

	return total
}

func checkMatches(s scratchcard) int {
	matches := 0
	for _, w := range s.winners {
		for _, n := range s.numbers {
			if w == n {
				matches += 1
			}
		}
	}
	return matches
}

func readScratchcards(file string) []scratchcard {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read %v: %v", file, err)
	}

	lines := strings.Split(trim(string(f)), "\n")

	returnScratchcards := []scratchcard{}

	for i, v := range lines {
		returnScratchcards = append(returnScratchcards, parseCard(v))
		returnScratchcards[i].count = 1
	}

	return returnScratchcards
}

func parseCard(line string) scratchcard {
	winnersAndNumbers := strings.Split(strings.Split(line, ":")[1], "|")
	re := regexp.MustCompile(`\s+`)

	winners, numbers := re.Split(trim(winnersAndNumbers[0]), -1), re.Split(trim(winnersAndNumbers[1]), -1)
	returnScratchcard := scratchcard{}

	for _, v := range winners {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Could not convert %v: %v", v, err)
		}
		returnScratchcard.winners = append(returnScratchcard.winners, s)
	}

	for _, v := range numbers {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Could not convert %v: %v", v, err)
		}
		returnScratchcard.numbers = append(returnScratchcard.numbers, s)
	}

	return returnScratchcard
}

func trim(line string) string {
	return strings.TrimSpace(line)
}
