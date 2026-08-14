package main

import (
	"fmt"
	"log"
	"os"
)

func getPuzzle(day string) string {
	content, err := os.ReadFile(day)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func firstPart() string {
	level := 0
	puzzle := getPuzzle("day1.txt")
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
	puzzle := getPuzzle("day1.txt")
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
