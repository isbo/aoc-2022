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
	moveMap := map[string]string{
		"A X": "Z",
		"A Y": "X",
		"A Z": "Y",
		"B X": "X",
		"B Y": "Y",
		"B Z": "Z",
		"C X": "Y",
		"C Y": "Z",
		"C Z": "X",
	}
	file, e := os.Open("input")
	utils.PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var score = 0
	for scanner.Scan() {
		line := scanner.Text()
		theirMove := strings.Split(line, " ")[0]
		myMove := moveMap[line]
		score += outcomeScore[theirMove+" "+myMove]
		score += shapeScore[myMove]
	}
	fmt.Println("Score", score)

}
