package main

import (
	. "aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getXValues(insts []string) []int {
	vals := make([]int, 242) // hardcode the max value we'll see
	var cycle = 1
	var x = 1
	vals[cycle] = x
	for _, inst := range insts {
		if inst == "noop" {
			cycle++
			vals[cycle] = x
		} else {
			inc, _ := strconv.Atoi(inst[5:])
			vals[cycle+1] = x
			x += inc
			vals[cycle+2] = x
			cycle += 2
		}
	}
	return vals
}

func part1(insts []string) int {
	vals := getXValues(insts)
	var ss = 0
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		ss += c * vals[c+1]
	}
	return ss
}

func part2(insts []string) {
	vals := getXValues(insts)
	var cycle = 1
	for r := 0; r < 6; r++ {
		for p := 0; p < 40; p++ {
			x := vals[cycle]
			if p >= x-1 && p <= x+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			cycle++
		}
		fmt.Println()
	}
}

func main() {
	file, e := os.Open("input")
	PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("part1", part1(lines))
	part2(lines)
}
