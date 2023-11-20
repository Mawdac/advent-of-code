package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	itsHappening := false

	data := readFile("input.txt")

	// for each number
	for i, v1 := range data {
		fmt.Println("Index:", i)
		fmt.Println("Value:", v1)
		for _, v2 := range data[i+1:] {
			for _, v3 := range data[i+2:] {
				itsHappening = checkFor2020(v1, v2, v3)
				if itsHappening {
					fmt.Println(v1, "plus", v2, "plus", v3, "is 2020!")
					fmt.Println("Multipled they are:", v1*v2*v3)
					break
				}
			}
			if itsHappening {
				break
			}
		}
		if itsHappening {
			break
		}
	}
}

func checkFor2020(i1 int, i2 int, i3 int) bool {
	return i1+i2+i3 == 2020
}

func readFile(file string) []int {
	intSlice := make([]int, 0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting data:", err)
			os.Exit(1)
		}
		intSlice = append(intSlice, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return intSlice
}
