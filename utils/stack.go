package utils

type Stack[T any] []T

func (s Stack[T]) Push(val T) Stack[T] {
	return append(s, val)
}

func (s Stack[T]) Pop() (Stack[T], T) {
	return s[:len(s)-1], s[len(s)-1]
}

func (s Stack[T]) Peek() T {
	return s[len(s)-1]
}

func (s Stack[T]) Empty() bool {
	return len(s) == 0
}
