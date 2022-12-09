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

type knot struct {
	x, y int
}

func (k *knot) move(x int, y int) {
	k.x, k.y = x, y
}

func (k *knot) follow(fk *knot) {
	if !k.isDisconnected(fk) {
		return
	}
	if fk.x > k.x {
		k.x += 1
	} else if fk.x < k.x {
		k.x -= 1
	}
	if fk.y > k.y {
		k.y += 1
	} else if fk.y < k.y {
		k.y -= 1
	}
}

func (k *knot) isDisconnected(fk *knot) bool {
	return math.Abs(float64(k.x-fk.x))+math.Abs(float64(k.y-fk.y)) > 1 &&
		(math.Abs(float64(k.x-fk.x)) != 1 || math.Abs(float64(k.y-fk.y)) != 1)
}

func countVisitsPart2(motions []string, n int) int {
	v := make(map[string]bool)
	v["0 0"] = true
	knots := make([]*knot, n)
	for i := range knots {
		knots[i] = new(knot)
	}
	h := knots[0]
	t := knots[n-1]
	for _, motion := range motions {
		f := strings.Split(motion, " ")
		steps, _ := strconv.Atoi(f[1])
		for i := 0; i < steps; i++ {
			var dx, dy int
			switch f[0] {
			case "R":
				dx = 1
			case "L":
				dx = -1
			case "U":
				dy = 1
			case "D":
				dy = -1
			}
			h.move(h.x+dx, h.y+dy)

			for i := 1; i < n; i++ {
				knots[i].follow(knots[i-1])
			}
			v[fmt.Sprintf("%d %d", t.x, t.y)] = true
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

	fmt.Println("visited p1", countVisits(lines))
	fmt.Println("visited p2", countVisitsPart2(lines, 10))
}
