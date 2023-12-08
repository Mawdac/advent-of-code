package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type desertMap struct {
	directions []rune
	nodes      []node
}

type node struct {
	left  *node
	right *node
	value string
}

func main() {
	fmt.Println("Hello 2023 day 8!")
	theMap := readDesertMap("input.txt")
	theSecondMap := readDesertMap("input.txt")
	// fmt.Println(string(theMap.directions))

	fmt.Printf("The solution to part 1 is %v\n", part1(theMap))
	before := time.Now()
	fmt.Printf("The solution to part 2 is %v\n", part2(theSecondMap))
	afterp2 := time.Now()
	fmt.Printf("It took %v to generate part 2\n", afterp2.Sub(before))
}

func part1(m desertMap) int {
	// start at 'AAA'
	var location *node
	jumps := 0
	for _, n := range m.nodes {
		if n.value == "AAA" {
			location = &n
			break
		}
	}

	fmt.Println("Starting at AAA...")
	// go through the directions forever
	for i := 0; true; i++ {
		if i == len(m.directions) {
			i = 0
		}

		fmt.Println("We're at", location.value)
		if m.directions[i] == 'L' {
			location = location.left
			fmt.Println("Going left to", location.value)
		} else {
			location = location.right
			fmt.Println("Going right to", location.value)
		}
		jumps++

		// stop if we're at 'ZZZ'
		if location.value == "ZZZ" {
			break
		}
	}
	return jumps
}

func part2(m desertMap) int {
	var locations []node
	jumps := 0
	// for every node
	for _, n := range m.nodes {
		// if it ends with 'A', add it to the location list
		if n.value[len(n.value)-1] == 'A' {
			fmt.Println("Adding", n)
			locations = append(locations, n)
		}
	}

	for i, v := range locations {
		fmt.Printf("Location %v: %v\n", i, v.value)
	}

	// go through directions forever
	for i := 0; true; i++ {
		if i == len(m.directions) {
			i = 0
		}

		// fmt.Println("Going", string(m.directions[i]))
		for j := 0; j < len(locations); j++ {
			// fmt.Println("We're at", locations[j].value)
			if m.directions[i] == 'L' {
				locations[j] = *locations[j].left
				// fmt.Println("Going left to", locations[j].value)
			} else {
				locations[j] = *locations[j].right
				// fmt.Println("Going right to", locations[j].value)
			}
		}
		jumps++
		// for i, v := range locations {
		// 	fmt.Printf("After Locations %v: %v\n", i, v.value)
		// }

		// stop if all locations end in Z
		if checkEnd(locations) {
			break
		}
	}
	return jumps
}

func checkEnd(nodes []node) bool {
	for _, n := range nodes {
		if rune(n.value[len(n.value)-1]) != 'Z' {
			return false
		}
	}
	return true
}

func readDesertMap(file string) desertMap {
	returnDesertMap := desertMap{}
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v\n", file, err)
	}

	directionsAndNodes := strings.Split(trim(string(f)), "\n\n")

	returnDesertMap.directions = []rune(directionsAndNodes[0])

	nodeLines := strings.Split(directionsAndNodes[1], "\n")

	tempStringsToNodes := map[string]*node{}
	tempStringsToStrings := make(map[string][]string, 0)

	// initialize the value of each node, and fill up temp objects so that we can
	// initialize the node pointers on the second pass
	for _, l := range nodeLines {
		valueAndNodes := strings.Split(l, "=")
		value := trim(string(valueAndNodes[0]))
		tempStringsToNodes[value] = &node{value: value} // we'll use this to setup the node pointers

		leftAndRight := strings.Split(valueAndNodes[1], ",")
		leftString := string([]rune(trim(leftAndRight[0]))[1:])
		r := len([]rune(trim(leftAndRight[1]))) - 1
		rightString := string([]rune(trim(leftAndRight[1]))[:r])

		tempStringsToStrings[value] = []string{leftString, rightString} // and this
	}

	for v, n := range tempStringsToStrings {
		fmt.Printf("Value: %v, Left %v, Right %v*\n", v, n[0], n[1])
		tempStringsToNodes[v].left = tempStringsToNodes[n[0]]
		tempStringsToNodes[v].right = tempStringsToNodes[n[1]]
	}

	for _, v := range tempStringsToNodes {
		returnDesertMap.nodes = append(returnDesertMap.nodes, *v)
	}

	return returnDesertMap
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
