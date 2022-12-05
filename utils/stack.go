package utils

type Stack []byte

func (s Stack) Push(b byte) Stack {
	return append(s, b)
}

func (s Stack) Pop() (Stack, byte) {
	return s[:len(s)-1], s[len(s)-1]
}
