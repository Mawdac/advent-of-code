package main

import "fmt"

func main() {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	fmt.Printf("Total increases on sweep depth: %d\n", measureSweeps(input))
}

func measureSweeps(sweeps []int) int {
	var previous, increases int

	for i, v := range sweeps {
		if i == 0 {
			fmt.Printf("Skipping first iteration %d\n", v)
			previous = v
			continue
		} else if v > previous {
			increases++
		}
		fmt.Printf("Increases: %d - Previous: %d - Current %d\n", increases, previous, v)
		previous = v
	}

	return increases
}
