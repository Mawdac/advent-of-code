package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type almanac struct {
	seeds []int
	maps  []conversionMap
}
type conversionMap struct {
	from     string
	to       string
	formulae []formula
}
type formula struct {
	sourceRangeStart      int
	destinationRangeStart int
	lengthRange           int
}

func main() {
	before := time.Now()
	fmt.Println("Hello, 2023 day 5!")
	almightyAlmanac := readAlmanac("input.txt")

	fmt.Printf("The seeds: %v\n", almightyAlmanac.seeds)

	for i, m := range almightyAlmanac.maps {
		fmt.Printf("\nMap %v\n\nFrom: %v\t To:%v\n", i, m.from, m.to)
		for _, f := range m.formulae {
			fmt.Printf(
				"Source: %v\t Dest: %v\t Range: %v\n",
				f.sourceRangeStart,
				f.destinationRangeStart,
				f.lengthRange,
			)
		}
	}

	fmt.Printf("The solution to part 1 is %v\n", part1(almightyAlmanac))
	afterp1 := time.Now()
	fmt.Printf("It took %v to generate\n", afterp1.Sub(before))
	before = time.Now()
	fmt.Printf("The solution to part 2 is %v\n", part2(almightyAlmanac))
	afterp2 := time.Now()
	fmt.Printf("It took %v to generate\n", afterp2.Sub(before))
}

func part1(allie almanac) int {
	lowestLocation := 0
	for _, s := range allie.seeds {
		seedLocation := getSeedLocation(s, allie.maps)
		if lowestLocation == 0 || seedLocation < lowestLocation {
			fmt.Printf("New lowest %v\n", seedLocation)
			lowestLocation = seedLocation
		}
	}
	return lowestLocation
}

func part2(allie almanac) int {
	lowestLocation := 0
	for i := 0; i < len(allie.seeds); i += 2 {
		for j := allie.seeds[i]; j < allie.seeds[i]+allie.seeds[i+1]; j++ {
			seedLocation := getSeedLocation(j, allie.maps)
			if lowestLocation == 0 || seedLocation < lowestLocation {
				fmt.Printf("New lowest %v\n", seedLocation)
				lowestLocation = seedLocation
			}
		}
	}
	return lowestLocation
}

func getSeedLocation(seed int, maps []conversionMap) int {
	location := seed
	for _, v := range maps {
		location = convertNumber(location, v)
	}
	return location
}

func convertNumber(n int, m conversionMap) int {
	for _, f := range m.formulae {
		if f.sourceRangeStart <= n && n < f.sourceRangeStart+f.lengthRange {
			return f.destinationRangeStart + (n - f.sourceRangeStart)
		}
	}
	return n
}

func readAlmanac(file string) almanac {
	returnAlmanac := almanac{}
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("We couldn't read %v, %v\nDo better.\n", file, err)
	}

	// split by double newlines
	sections := strings.Split(trim(string(f)), "\n\n")

	// seeds
	seeds := strings.Split(trim(strings.Split(sections[0], ":")[1]), " ")
	for _, v := range seeds {
		s, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Couldn't convert seed %v, %v\n", v, err)
		}
		returnAlmanac.seeds = append(returnAlmanac.seeds, s)
	}

	// conversionMaps
	for _, m := range sections[1:] {
		returnMap := conversionMap{}
		lines := strings.Split(m, "\n")
		fromAndTo := strings.Split(strings.Split(trim(lines[0]), " ")[0], "-")
		returnMap.from, returnMap.to = fromAndTo[0], fromAndTo[2]
		for _, f := range lines[1:] {
			numbers := strings.Split(f, " ")
			destinationRange, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatalf("Couldn't parse dest range %v, %v\n", numbers[0], err)
			}
			sourceRange, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatalf("Couldn't parse source range %v, %v\n", numbers[1], err)
			}
			length, err := strconv.Atoi(numbers[2])
			if err != nil {
				log.Fatalf("Couldn't parse range length %v, %v\n", numbers[2], err)
			}
			returnMap.formulae = append(returnMap.formulae, formula{
				destinationRangeStart: destinationRange,
				sourceRangeStart:      sourceRange,
				lengthRange:           length,
			})
		}
		returnAlmanac.maps = append(returnAlmanac.maps, returnMap)
	}

	return returnAlmanac
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
