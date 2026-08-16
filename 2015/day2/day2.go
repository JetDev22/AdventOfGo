package main

import (
	"AdventOfGo/helper"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func firstPart() float64 {
	// lwh formula = 2*l*w + 2*w*h + 2*h*l
	total := 0.0
	puzzle := helper.GetPuzzle("day2.txt")
	seperated := strings.Split(puzzle, "\n")
	for _, single := range seperated {
		present := strings.Split(single, "x")
		l, _ := strconv.ParseFloat(present[0], 64)
		w, _ := strconv.ParseFloat(present[1], 64)
		h, _ := strconv.ParseFloat(present[2], 64)
		sides := []float64{l, w, h}
		slices.Sort(sides)
		sideA := 2 * l * w
		sideB := 2 * w * h
		sideC := 2 * h * l
		total += sideA + sideB + sideC + (sides[0] * sides[1])
	}
	return total
}

func main() {
	fmt.Println(firstPart())
	//fmt.Println(secondPart())
}
