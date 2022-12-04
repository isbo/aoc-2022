package main

import (
	"aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, e := os.Open("input")
	utils.PanicOnError(e)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`[-,]`)
	var subset = 0
	for scanner.Scan() {
		line := scanner.Text()
		f := re.Split(line, -1)
		if overlaps(f) {
			subset += 1
		}
	}
	fmt.Println("contained ", subset)
}

func overlaps(f []string) bool {
	var n [4]int
	for i, e := range f {
		v, _ := strconv.Atoi(e)
		n[i] = v
	}
	return (n[1] >= n[2] && n[1] <= n[3]) ||
		(n[0] >= n[2] && n[0] <= n[3]) ||
		(n[0] <= n[2] && n[3] <= n[1]) ||
		(n[0] >= n[2] && n[1] <= n[3])
}
