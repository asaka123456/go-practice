package main

import "fmt"

type stack struct {
	s    []any
	head int
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

func push(s *stack, v any) {
	if s.head >= len(s.s)-1 {
		fmt.Println("Стек полон")
		return
	}
	s.head++
	s.s[s.head] = v
}

func pop(s *stack) any {
	if s.head < 0 {
		fmt.Println("Стек пуст")
		return nil
	}
	val := s.s[s.head]
	s.head--
	return val
}

func peek(s *stack) any {
	if s.head < 0 {
		fmt.Println("Стек пуст")
		return nil
	}
	return s.s[s.head]
}

func main() {
	s := newStack(5)
	push(s, 10)
	push(s, 20)
	push(s, 30)
	fmt.Println("Вершина:", peek(s))
	fmt.Println("Извлечено:", pop(s))
	fmt.Println("Извлечено:", pop(s))
	fmt.Println("Извлечено:", pop(s))
	fmt.Println("Извлечено:", pop(s))
}
