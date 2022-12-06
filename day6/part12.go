package main

import (
	. "aoc-2022/utils"
	"fmt"
	"os"
)

func getMarker(line string, length int) int {
	var sIdx = 0
	var mIdx = -1
	for i := 0; i < len(line); i++ {
		for j := sIdx; j < i; j++ {
			if line[i] == line[j] {
				sIdx = j + 1
				break
			}
		}
		if i-sIdx >= length-1 {
			mIdx = i + 1
			break
		}
	}
	return mIdx
}

func main() {
	b, e := os.ReadFile("input")
	PanicOnError(e)

	line := string(b)
	fmt.Println("marker indices: ", getMarker(line, 4), getMarker(line, 14))
}
