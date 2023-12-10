package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type tile struct {
	shape    rune
	outline  bool
	enclosed bool
	x        int
	y        int
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
	fmt.Printf("Tiles enclosed in the pipe: %v\n", part2(tiles))
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

func part2(m [][]tile) int {
	// find 'S' to start
	location := findStart(m)
	m[location.y][location.x].outline = true
	previousLocation := ""

	for {
		fmt.Printf("Currently at %v\n", string(location.shape))
		fmt.Printf("Previous location: %v\n", previousLocation)
		location, previousLocation = findNextLocation(m, location, previousLocation)
		m[location.y][location.x].outline = true
		if location.shape == 'S' {
			break
		}
	}

	for _, v := range m {
		fmt.Println()
		for _, j := range v {
			if j.outline {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()

	enclosedTiles := 0
	for y, r := range m {
		fmt.Println()
		for x, t := range r {
			if t.outline {
				fmt.Print("|")
			} else {
				if checkEnclosed(m, t) {
					fmt.Print("+")
					m[y][x].enclosed = true
				} else {
					fmt.Print(".")
				}
			}
		}
	}

	cleanupEnclosed(m) // cleans up some enclosed tiles that were false positives

	for _, r := range m {
		fmt.Println()
		for _, t := range r {
			if t.outline {
				fmt.Print("|")
			} else if t.enclosed {
				enclosedTiles++
				fmt.Print("+")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()

	return enclosedTiles
}

func cleanupEnclosed(m [][]tile) int {
	enclosedTilesCleaned := 0
	for y, r := range m {
		for x, t := range r {
			if t.enclosed {
				// north
				if y-1 > 0 && !m[y-1][x].enclosed && !m[y-1][x].outline {
					m[y][x].enclosed = false
					enclosedTilesCleaned++
					continue
				}
				// east
				if x+1 < len(m[y])-1 && !m[y][x+1].enclosed && !m[y][x+1].outline {
					m[y][x].enclosed = false
					enclosedTilesCleaned++
					continue
				}
				// south
				if y+1 < len(m)-1 && !m[y+1][x].enclosed && !m[y+1][x].outline {
					m[y][x].enclosed = false
					enclosedTilesCleaned++
					continue
				}
				// east
				if x-1 > 0 && !m[y][x-1].enclosed && !m[y][x-1].outline {
					m[y][x].enclosed = false
					enclosedTilesCleaned++
					continue
				}
			}
		}
	}
	if enclosedTilesCleaned != 0 {
		fmt.Printf("Cleaned up %v tiles, taking another pass...\n", enclosedTilesCleaned)
		return cleanupEnclosed(m)
	} else {
		fmt.Println("No more tiles to cleanup! 😃")
		return 0
	}
}

func checkEnclosed(m [][]tile, t tile) bool {
	n, e, s, w := false, false, false, false
	// north
	for i := t.y - 1; i > -1; i-- {
		if m[i][t.x].outline {
			n = true
		}
	}
	// east
	for i := t.x + 1; i < len(m[t.y])-1; i++ {
		if m[t.y][i].outline {
			e = true
		}
	}
	// south
	for i := t.y + 1; i < len(m)-1; i++ {
		if m[i][t.x].outline {
			s = true
		}
	}
	// west
	for i := t.x - 1; i > -1; i-- {
		if m[t.y][i].outline {
			w = true
		}
	}
	return n && e && s && w
}

func findNextLocation(m [][]tile, t tile, previous string) (tile, string) {
	fmt.Printf("Checking next location for %v: %v, %v...\n", string(t.shape), t.y, t.x)
	// north
	if t.y-1 > -1 && compass[t.shape].north && compass[m[t.y-1][t.x].shape].south && previous != "north" {
		return m[t.y-1][t.x], "south"
	} // east
	if t.x+1 < len(m[t.y]) && compass[t.shape].east && compass[m[t.y][t.x+1].shape].west && previous != "east" {
		return m[t.y][t.x+1], "west"
	} // south
	if t.y+1 < len(m) && compass[t.shape].south && compass[m[t.y+1][t.x].shape].north && previous != "south" {
		return m[t.y+1][t.x], "north"
	} // west
	if t.x-1 > -1 && compass[t.shape].west && compass[m[t.y][t.x-1].shape].east && previous != "west" {
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
			returnRow = append(returnRow, tile{
				shape:    t,
				outline:  false,
				enclosed: false,
				x:        x,
				y:        y,
			})
		}
		returnTiles = append(returnTiles, returnRow)
	}
	return returnTiles
}

func trim(text string) string {
	return strings.TrimSpace(text)
}
