package main

import (
	"AdventOfGo/helper"
	"fmt"
)

func firstPart() string {
	level := 0
	puzzle := helper.GetPuzzle("day1.txt")
	for _, char := range puzzle {
		if char == '(' {
			level++
		} else if char == ')' {
			level--
		}
	}
	return fmt.Sprintf("Solution for Part 1 = %d", level)
}

func secondPart() string {
	level := 0
	firstBase := 0
	puzzle := helper.GetPuzzle("day1.txt")
	for i, char := range puzzle {
		if char == '(' {
			level++
		} else if char == ')' {
			level--
		}
		if level == -1 {
			firstBase = i + 1
			return fmt.Sprintf("Solution for Part 2 = %d", firstBase)
		}
	}
	return ""
}

func main() {
	fmt.Println(firstPart())
	fmt.Println(secondPart())
}
