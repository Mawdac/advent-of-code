package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time   int
	record int
}

func main() {
	fmt.Println("Hello, 2023 day 6!")
	races := parseRaces("input.txt")

	for i, r := range races {
		fmt.Printf("Race #%v\ttime:%v\tdistance:%v\n", i, r.time, r.record)
	}
	fmt.Printf("The solution to part 1 is %v\n", part1(races))
	fmt.Printf("The solution to part 2 is %v\n", part2(races))
}

func part1(races []race) int {
	returnTotal := 1
	for _, r := range races {
		returnTotal *= checkRace(r)
	}
	return returnTotal
}

func part2(races []race) int {
	timeString, recordString := "", ""
	for _, r := range races {
		timeString += fmt.Sprint(r.time)
		recordString += fmt.Sprint(r.record)
	}

	newTime, _ := strconv.Atoi(timeString)
	newRecord, _ := strconv.Atoi(recordString)

	return checkRace(race{time: newTime, record: newRecord})
}

func checkRace(r race) int {
	waysToWin := 0
	for i := 1; i < r.time; i++ {
		distance := i * (r.time - i)
		if distance > r.record {
			waysToWin++
		}
	}
	return waysToWin
}

func parseRaces(file string) []race {
	returnRaces := make([]race, 0)
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v", file, err)
	}

	lines := strings.Split(trim(string(f)), "\n")
	re := regexp.MustCompile(`\s+`)

	times, distances := re.Split(lines[0], -1), re.Split(lines[1], -1)

	for i := 1; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatalf("Couldn't parse time %v, %v", times[i], err)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatalf("Couldn't parse time %v, %v", distances[i], err)
		}
		returnRaces = append(returnRaces, race{time: time, record: distance})
	}

	return returnRaces
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
