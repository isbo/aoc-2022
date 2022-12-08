package utils

type Stack[T any] struct {
	vals []T
}

func (stack *Stack[T]) Push(val T) {
	stack.vals = append(stack.vals, val)
}

func (stack *Stack[T]) Pop() T {
	topVal := stack.vals[(len(stack.vals) - 1)]
	stack.vals = stack.vals[:len(stack.vals)-1]
	return topVal
}

func (stack *Stack[T]) Peek() T {
	return stack.vals[len(stack.vals)-1]
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.vals) == 0
}
