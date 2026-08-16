package helper

import (
	"log"
	"os"
)

func GetPuzzle(day string) string {
	content, err := os.ReadFile(day)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
