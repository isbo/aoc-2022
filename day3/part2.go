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
	var i = 0
	var groups [3]string
	for scanner.Scan() {
		ks := scanner.Text()
		groups[i] = ks
		if i < 2 {
			i += 1
			continue
		}
		i = 0
		m1 := make(map[rune]bool)
		m2 := make(map[rune]bool)
		for _, c := range groups[0] {
			m1[c] = true
		}
		for _, c := range groups[1] {
			if _, exists := m1[c]; exists {
				m2[c] = true
			}
		}
		for _, c := range groups[2] {
			if _, exists := m2[c]; exists {
				priority += scoreMap[c]
				break
			}
		}
	}
	fmt.Println("priority ", priority)
}
