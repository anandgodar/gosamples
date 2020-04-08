package main

import "fmt"

func main() {
	s1 := "vista" // try changing "fish" for "vista"
	s2 := "hish"

	grid := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		grid[i] = make([]int, len(s2))
	}

	var maxValue int
	for i, s1_char := range s1 {
		for j, s2_char := range s2 {
			// fmt.Printf("s1_char: %c, s1_char: %c\n", s1_char, s2_char) // nice debug
			if s1_char != s2_char {
				grid[i][j] = 0
			} else {
				// verify if grid position is valid
				if i-1 >= 0 && j-i >= 0 {
					maxValue = grid[i-1][j-1] + 1
					grid[i][j] = maxValue
				}
			}
		}
	}

	fmt.Printf("Longest common substring size between %q and %q is %d!\n", s1, s2, maxValue)
}