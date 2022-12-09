package main

import (
	. "aoc-2022/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func countVisits(motions []string) int {
	var tx, ty, hx, hy, ohx, ohy int // graph-paper coordinates

	v := make(map[string]bool)
	v["0 0"] = true
	for _, motion := range motions {
		f := strings.Split(motion, " ")
		steps, _ := strconv.Atoi(f[1])
		for i := 0; i < steps; i++ {
			ohx, ohy = hx, hy
			switch f[0] {
			case "R":
				hx++
			case "L":
				hx--
			case "U":
				hy++
			case "D":
				hy--
			}

			if math.Abs(float64(hx-tx))+math.Abs(float64(hy-ty)) > 1 &&
				(math.Abs(float64(hx-tx)) != 1 || math.Abs(float64(hy-ty)) != 1) {
				tx = ohx
				ty = ohy
				v[fmt.Sprintf("%d %d", tx, ty)] = true
			}
		}
	}
	return len(v)
}

func main() {
	file, e := os.Open("input")
	PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("visited", countVisits(lines))
}
