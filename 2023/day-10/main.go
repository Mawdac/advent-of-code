package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type tile struct {
	shape rune
	x     int
	y     int
}

type pipe struct {
	north bool
	east  bool
	south bool
	west  bool
}

var compass = map[rune]pipe{
	'|': {north: true, east: false, south: true, west: false},
	'-': {north: false, east: true, south: false, west: true},
	'L': {north: true, east: true, south: false, west: false},
	'J': {north: true, east: false, south: false, west: true},
	'7': {north: false, east: false, south: true, west: true},
	'F': {north: false, east: true, south: true, west: false},
	'.': {north: false, east: false, south: false, west: false},
	'S': {north: true, east: true, south: true, west: true},
}

func main() {
	fmt.Println("Hello, 2023 day 10!")

	tiles := readTiles("input.txt")
	for _, r := range tiles {
		for _, t := range r {
			fmt.Print(string(t.shape))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Steps to the farthest position: %v\n", part1(tiles))
}

func part1(m [][]tile) int {
	// find 'S' to start
	location := findStart(m)
	previousLocation := ""
	hops := 0

	for {
		fmt.Printf("Currently at %v\n", string(location.shape))
		fmt.Printf("Previous location: %v\n", previousLocation)
		location, previousLocation = findNextLocation(m, location, previousLocation)
		hops++
		if location.shape == 'S' {
			break
		}
	}
	fmt.Printf("Total length of pipe: %v\n", hops)
	return hops / 2
}

func findNextLocation(m [][]tile, t tile, previous string) (tile, string) {
	// north
	if compass[t.shape].north && compass[m[t.y-1][t.x].shape].south && previous != "north" {
		return m[t.y-1][t.x], "south"
	} // east
	if compass[t.shape].east && compass[m[t.y][t.x+1].shape].west && previous != "east" {
		return m[t.y][t.x+1], "west"
	} // south
	if compass[t.shape].south && compass[m[t.y+1][t.x].shape].north && previous != "south" {
		return m[t.y+1][t.x], "north"
	} // west
	if compass[t.shape].west && compass[m[t.y][t.x-1].shape].east && previous != "west" {
		return m[t.y][t.x-1], "east"
	}
	panic("Couldn't find the next tile!")
}

func findStart(m [][]tile) tile {
	for _, r := range m {
		for _, t := range r {
			if t.shape == 'S' {
				return t
			}
		}
	}
	panic("Couldnt' find starting location!")
}

func readTiles(file string) [][]tile {
	returnTiles := [][]tile{}
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read %v, %v\n", file, err)
	}
	lines := strings.Split(trim(string(f)), "\n")

	for y, line := range lines {
		returnRow := []tile{}
		for x, t := range line {
			returnRow = append(returnRow, tile{shape: t, x: x, y: y})
		}
		returnTiles = append(returnTiles, returnRow)
	}
	return returnTiles
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
