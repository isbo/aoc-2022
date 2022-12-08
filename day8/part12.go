package main

import (
	. "aoc-2022/utils"
	"bufio"
	"fmt"
	"os"
)

// this is linear in the size of input (times 2 because of two passes per row/column)
func countVisible(t [][]uint8) int {
	h := len(t)
	w := len(t[0])

	var count = 0
	v := make(map[int]bool)
	// in each row, find number of visible trees from left and right
	for r := 1; r < h-1; r++ {
		var max = t[r][0]
		for c := 1; c < w-1; c++ {
			if t[r][c] > max {
				max = t[r][c]
				v[r*w+c] = true
			}
		}
		max = t[r][w-1]
		for c := w - 2; c > 0; c-- {
			if t[r][c] > max {
				max = t[r][c]
				v[r*w+c] = true
			}
		}
	}
	// in each column, find number of visible trees from top and bottom
	for c := 1; c < w-1; c++ {
		var max = t[0][c]
		for r := 1; r < h-1; r++ {
			if t[r][c] > max {
				max = t[r][c]
				v[r*w+c] = true
			}
		}
		max = t[h-1][c]
		for r := h - 2; r > 0; r-- {
			if t[r][c] > max {
				max = t[r][c]
				v[r*w+c] = true
			}
		}
	}
	return count + 2*w + 2*h - 4 + len(v)
}

// a linear time algorithm that is kinda complicated to implement
func findMostScenic(t [][]uint8) int {
	h := len(t)
	w := len(t[0])

	d := make([][]int, h)
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		for j := range d[i] {
			d[i][j] = 1
		}
	}
	// in each row, compute left and right viewing distances
	for r := 1; r < h-1; r++ {
		// as we scan, we keep track of the potential 'horizon' trees in the stack
		// if current tree is shorter than previous:
		//   we store it because it might be the 'horizon' for some future tree
		//   its viewing distance will be 1
		// if current tree is taller than previous:
		//	 we store it but also remove all previous 'horizon' trees that are shorter than it
		//   because all those trees cannot be a 'horizon' for any future tree shorter than the current one
		//   the viewing distance would be the distance to the next taller horizon tree in the stack
		// each row/column scan is O(n) where n is row/column size
		// total of adds/removals from stack is O(n) because an element can be added or removed at most one time.

		var s = new(Stack[int])
		s.Push(0)
		for c := 1; c < w-1; c++ {
			var dist = 1
			if t[r][c] > t[r][c-1] {
				for !s.Empty() && t[r][s.Peek()] < t[r][c] {
					s.Pop()
				}
				var i = 0
				if !s.Empty() {
					i = s.Peek()
				}
				dist = c - i
			}
			s.Push(c)
			d[r][c] *= dist
		}
		s = new(Stack[int])
		s.Push(w - 1)
		for c := w - 2; c > 0; c-- {
			var dist = 1
			if t[r][c] > t[r][c+1] {
				for !s.Empty() && t[r][s.Peek()] < t[r][c] {
					s.Pop()
				}
				var i = w - 1
				if !s.Empty() {
					i = s.Peek()
				}
				dist = i - c
			}
			s.Push(c)
			d[r][c] *= dist
		}
	}
	// in each row, compute top and bottom viewing distances
	for c := 1; c < w-1; c++ {
		var s = new(Stack[int])
		s.Push(0)
		for r := 1; r < h-1; r++ {
			var dist = 1
			if t[r][c] > t[r-1][c] {
				for !s.Empty() && t[s.Peek()][c] < t[r][c] {
					s.Pop()
				}
				var i = 0
				if !s.Empty() {
					i = s.Peek()
				}
				dist = r - i
			}
			s.Push(r)
			d[r][c] *= dist
		}

		s = new(Stack[int])
		s.Push(h - 1)
		for r := h - 2; r > 0; r-- {
			var dist = 1
			if t[r][c] > t[r+1][c] {
				for !s.Empty() && t[s.Peek()][c] < t[r][c] {
					s.Pop()
				}
				var i = h - 1
				if !s.Empty() {
					i = s.Peek()
				}
				dist = i - r
			}
			s.Push(r)
			d[r][c] *= dist
		}
	}
	var max = 0
	for i := range d {
		for j := range d[i] {
			if d[j][i] > max {
				max = d[j][i]
			}
		}
	}
	return max
}

func main() {
	file, e := os.Open("input")
	PanicOnError(e)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	t := make([][]uint8, len(lines))
	for r, line := range lines {
		t[r] = make([]uint8, len(line))
		for c := 0; c < len(line); c++ {
			t[r][c] = line[c] - 48
		}
	}
	fmt.Println("visible", countVisible(t))
	fmt.Println("scenic score", findMostScenic(t))
}
