package main

import (
	"aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	outcomeScore := map[string]int{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}
	shapeScore := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	file, e := os.Open("input")
	utils.PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var score = 0
	for scanner.Scan() {
		line := scanner.Text()
		answer := strings.Split(line, " ")[1]
		score += outcomeScore[line]
		score += shapeScore[answer]
	}
	fmt.Println("score", score)

}
