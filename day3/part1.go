package main

import (
	"aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, e := os.Open("input")
	utils.PanicOnError(e)

	scoreMap := make(map[rune]int)
	for i, c := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		scoreMap[c] = i + 1
	}

	scanner := bufio.NewScanner(file)
	var priority = 0
	for scanner.Scan() {
		ks := scanner.Text()
		m := make(map[rune]bool)
		for _, c := range ks[0 : len(ks)/2] {
			m[c] = true
		}
		for _, c := range ks[len(ks)/2:] {
			if _, exists := m[c]; exists {
				priority += scoreMap[c]
				break
			}
		}
	}
	fmt.Println("priority ", priority)

}
