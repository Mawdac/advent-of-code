package main

import (
	"testing"
)

func TestParseElvesFromFile(t *testing.T) {
	testFile := "test.txt"
	out := parseElvesFromFile(testFile)

	if len(out) != 5 {
		t.Errorf("Expected 5 elves, but got %v", len(out))
	}

	if out[0] != "1000\n2000\n3000" {
		t.Errorf(
			"Expected first elf to carry '1000\n2000\n3000', but got %v",
			out[0],
		)
	}

	if out[len(out)-1] != "10000" {
		t.Errorf(
			"Expected last elf to to carry '10000', but got %v",
			out[len(out)-1],
		)
	}
}

func TestParseCalories(t *testing.T) {
	testFile := "test.txt"
	elves := parseElvesFromFile(testFile)
	snacks := parseCalories(elves)

	if len(snacks) != 5 {
		t.Errorf("Expected 5 packs of snacks, but got %v", len(snacks))
	}

	if snacks[0] != 6000 {
		t.Errorf("Expected first elfs calories to be 6000, but got %v", snacks[0])
	}

	if snacks[len(snacks)-1] != 10000 {
		t.Errorf("Expected last elfs only pack to be 10000, but got %v", snacks[len(snacks)-1])
	}
}

func TestSortPacks(t *testing.T) {
	data := []int{1000, 3000, 8000, 4000, 5000}
	sorted := sortPacks(data)

	if sorted[0] != 8000 {
		t.Errorf("Expected most calories to be 8000, but got %v", sorted[0])
	}

	if sorted[len(data)-1] != 1000 {
		t.Errorf("Expected least calories to be 1000, but got %v", sorted[0])
	}
}

func TestPart1(t *testing.T) {
	data := []int{1000, 2000, 6969, 4000, 5000}
	top := part1(data)

	if top != 6969 {
		t.Errorf("Expected top elf to have 6969, but got %v", top)
	}
}

func TestPart2(t *testing.T) {
	data := []int{1000, 2000, 6969, 4000, 5000}
	top3 := part2(data)

	if len(top3) != 3 {
		t.Errorf("Expected top 3 elves, but got %v", len(top3))
	}

	if top3[0] != 6969 {
		t.Errorf("Expected top elf to have 6969, but got %v", top3[0])
	}
	if top3[1] != 5000 {
		t.Errorf("Expected second elf to have 5000, but got %v", top3[1])
	}
	if top3[2] != 4000 {
		t.Errorf("Expected third elf to have 4000, but got %v", top3[2])
	}
}
