package main

import (
	"aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, e := os.Open("input1")
	utils.PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var total = 0
	var cals [3]int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if total > cals[0] {
				cals[0] = total
				sort.Ints(cals[:])
			}
			total = 0
		} else {
			cals, e := strconv.Atoi(line)
			utils.PanicOnError(e)
			total += cals
		}
	}
	fmt.Println(cals[0] + cals[1] + cals[2])
}
