package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

func main() {
	fmt.Println("Hello 2021 day 2!")

	theFuckinCommands := readCommands("input.txt")

	for _, command := range theFuckinCommands {
		fmt.Printf("Direction: %v\t\tValue: %v\n", command.direction, command.value)
	}

	fmt.Printf("The solution to part 1 is %v\n", part1(theFuckinCommands))
	fmt.Printf("The solution to part 2 is %v\n", part2(theFuckinCommands))
}

func part1(commands []command) int {
	x, y := 0, 0

	// for every command
	for _, command := range commands {
		switch command.direction {
		case "forward":
			x += command.value
		case "down":
			y += command.value
		case "up":
			y += command.value * -1
		}
	}
	return x * y
}

func part2(commands []command) int {
	x, y, aim := 0, 0, 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			x += command.value
			y += aim * command.value
		case "down":
			aim += command.value
		case "up":
			aim += command.value * -1
		}
	}

	return x * y
}

func readCommands(file string) []command {
	returnCommands := []command{}
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v\n", file, err)
	}

	lines := strings.Split(trim(string(f)), "\n")

	for _, line := range lines {
		returnCommands = append(returnCommands, parseCommand(line))
	}

	return returnCommands
}

func parseCommand(line string) command {
	returnCommand := command{}
	directionAndValue := strings.Split(line, " ")

	direction := directionAndValue[0]
	valueString := directionAndValue[1]

	returnCommand.direction = direction

	value, err := strconv.Atoi(valueString)
	if err != nil {
		log.Fatalf("Couldn't convert %v, %v", valueString, err)
	}

	returnCommand.value = value

	return returnCommand
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
