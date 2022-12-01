package main

import (
	"aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, e := os.Open("input1")
	utils.PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var max, total = -1, 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if total > max {
				max = total
			}
			total = 0
		} else {
			cals, e := strconv.Atoi(line)
			utils.PanicOnError(e)
			total += cals
		}
	}
	fmt.Println(max)

}
