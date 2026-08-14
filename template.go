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

func main() {
	puzzle := getPuzzle("dayXpuzzle.txt")
	fmt.Println(puzzle)
}
