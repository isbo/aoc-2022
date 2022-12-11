package main

import (
	"fmt"
	"math"
	"sort"
)

type monkey struct {
	items     []int
	worryFn   func(int) int
	div       int
	tId       int
	fId       int
	inspected int
}

type itemMsg struct {
	worry  int
	monkey int
}

func (m *monkey) receive(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) playRound() []itemMsg {
	var msg []itemMsg
	for _, item := range m.items {
		m.inspected++
		worry := int(math.Floor(float64(m.worryFn(item) / 3.0)))
		var to int
		if worry%m.div == 0 {
			to = m.tId
		} else {
			to = m.fId
		}
		msg = append(msg, itemMsg{worry, to})
	}
	m.items = []int{}
	return msg
}

func add(inc int) func(int) int {
	return func(n int) int {
		return n + inc
	}
}

func multiply(inc int) func(int) int {
	return func(n int) int {
		return n * inc
	}
}

func square(val int) int {
	return val * val
}

func part1(monkeys []*monkey) int {
	for r := 0; r < 20; r++ {
		for _, m := range monkeys {
			messages := m.playRound()
			for _, msg := range messages {
				monkeys[msg.monkey].receive(msg.worry)
				//fmt.Printf("monkey %d passed %d to %d\n", i, msg.worry, msg.monkey)
			}
		}
		//for _, m := range monkeys {
		//	fmt.Println(r, m.items)
		//}
	}
	inspected := make([]int, len(monkeys))
	for i, m := range monkeys {
		//fmt.Println(i, m)
		inspected[i] = m.inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}

func main() {
	xMonkeys := []*monkey{
		{[]int{79, 98}, multiply(19), 23, 2, 3, 0},
		{[]int{54, 65, 75, 74}, add(6), 19, 2, 0, 0},
		{[]int{79, 60, 97}, square, 13, 1, 3, 0},
		{[]int{74}, add(3), 17, 0, 1, 0},
	}
	tMonkeys := []*monkey{
		{[]int{66, 71, 94}, multiply(5), 3, 7, 4, 0},
		{[]int{70}, add(6), 17, 3, 0, 0},
		{[]int{62, 68, 56, 65, 94, 78}, add(5), 2, 3, 1, 0},
		{[]int{89, 94, 94, 67}, add(2), 19, 7, 0, 0},
		{[]int{71, 61, 73, 65, 98, 98, 63}, multiply(7), 11, 5, 6, 0},
		{[]int{55, 62, 68, 61, 60}, add(7), 5, 2, 1, 0},
		{[]int{93, 91, 69, 64, 72, 89, 50, 71}, add(1), 13, 5, 2, 0},
		{[]int{76, 50}, square, 7, 4, 6, 0},
	}
	fmt.Println(tMonkeys)
	print("monkey business: ", part1(xMonkeys))
}
