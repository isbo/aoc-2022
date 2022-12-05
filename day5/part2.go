package main

import (
	. "aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseCrates(lines []string) []Stack {
	lastIdx := len(lines) - 1
	last := strings.TrimSpace(lines[lastIdx])
	n, _ := strconv.Atoi(last[strings.LastIndex(last, " ")+1:])
	crates := make([]Stack, n)

	for i := lastIdx - 1; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < n; j++ {
			crate := line[j*4+1]
			if crate != ' ' {
				crates[j] = crates[j].Push(crate)
			}
		}
	}
	return crates
}

func main() {
	file, e := os.Open("input")
	PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var lines []string
	var crates []Stack
	var parsed = false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 { // blank
			crates = parseCrates(lines)
			parsed = true
			continue
		} else if !parsed {
			lines = append(lines, line)
			continue
		}
		f := strings.Split(line, " ")
		num, _ := strconv.Atoi(f[1])
		from, _ := strconv.Atoi(f[3])
		to, _ := strconv.Atoi(f[5])
		var v byte
		var toMove []byte
		for i := 0; i < num; i++ {
			crates[from-1], v = crates[from-1].Pop()
			toMove = append(toMove, v)
		}
		for i := len(toMove) - 1; i >= 0; i-- {
			crates[to-1] = crates[to-1].Push(toMove[i])
		}
	}
	var out strings.Builder
	for _, crate := range crates {
		_, v := crate.Pop()
		out.WriteByte(v)
	}
	fmt.Println("top crates ", out.String())
}
