package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

type input struct {
	direction string
	distance  int
}

func newInput(direction string, distance int) input {
	i := input{direction: direction, distance: distance}
	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(file string) []input {
	// the final slice of directions we'll return
	var plans []input

	// open the file
	f, err := os.Open(file)
	check(err)

	// remember to close the file!
	defer f.Close()

	// returns a new scanner to read with
	scanner := bufio.NewScanner(f)

	// for each line
	for scanner.Scan() {
		// split the line by space
		textSlice := s.Split(scanner.Text(), " ")

		// get the dir and convert the distance to an int
		dir := textSlice[0]
		dis, err := strconv.Atoi(textSlice[1])
		check(err)

		// append the navigation plans with the new input
		plans = append(plans, newInput(dir, dis))

		// print the latest addition
		fmt.Printf("%+v\n", plans[len(plans)-1])
	}
	return plans
}

func main() {
	var forwardDistance, verticalDistance int
	plans := parseInput("input.txt")

	for range plans {
		switch i.direction {
		case condition:

		}
	}
}
