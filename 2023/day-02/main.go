package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}
type game struct {
	sets []set
}

func main() {
	var games []game
	if len(os.Args) > 1 {
		games = readGames(os.Args[1])
		fmt.Printf("Read in %v games", len(games))
	}

	if len(os.Args) > 2 {
		for i, v := range games {
			fmt.Println("Game: ", i+1)
			fmt.Println("Sets:", v)
		}
		if os.Args[2] == "part1" || os.Args[2] == "all" {
			p1 := part1(games)
			fmt.Println("~ Part 1 Solution ~")
			fmt.Printf("The solution for part 1 is %v\n", p1)
		}
		if os.Args[2] == "part2" || os.Args[2] == "all" {
			p2 := part2(games)
			fmt.Println("~ Part 2 Solution ~")
			fmt.Printf("The solution for part 2 is %v\n", p2)
		}
	} else {
		fmt.Println("Missing command flags")
		fmt.Println("Try this: 'go run main.go input.txt all/part1/part2'")
	}
}

func part1(games []game) int {
	total := 0
	for i, g := range games {
		possible := true
		for _, s := range g.sets {
			if s.red > 12 || s.green > 13 || s.blue > 14 {
				possible = false
			}
		}
		if possible {
			total += i + 1
		}
	}
	return total
}

func part2(games []game) int {
	total := 0
	for _, g := range games {
		var redMax, greenMax, blueMax int
		for _, s := range g.sets {
			if s.red > redMax {
				redMax = s.red
			}
			if s.green > greenMax {
				greenMax = s.green
			}
			if s.blue > blueMax {
				blueMax = s.blue
			}
		}
		total += (redMax * greenMax * blueMax)
	}
	return total
}

func readGames(file string) []game {
	returnGames := []game{}
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read %v: %v", file, err)
	}

	trimmedFile := strings.TrimSuffix(string(f), "\n")
	lines := strings.Split(trimmedFile, "\n")

	for _, v := range lines {
		returnGames = append(returnGames, parseGame(trim(v)))
	}

	return returnGames
}

func parseGame(line string) game {
	returnGame := game{}
	splitGameFromIndex := strings.Split(line, ":")
	splitSets := strings.Split(trim(splitGameFromIndex[1]), ";")
	for _, v := range splitSets {
		newSet := set{}
		splitCubes := strings.Split(trim(v), ",")
		for _, v2 := range splitCubes {
			countAndColor := strings.Split(trim(v2), " ")
			addColorCount(countAndColor, &newSet)
		}
		returnGame.sets = append(returnGame.sets, newSet)
	}
	return returnGame
}

func addColorCount(countAndColor []string, s *set) {
	count, _ := strconv.Atoi(countAndColor[0])
	switch countAndColor[1] {
	case "red":
		s.red += count
	case "green":
		s.green += count
	case "blue":
		s.blue += count
	}
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
