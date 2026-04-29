package main

import "fmt"

type queue struct {
	s    []any
	low  int
	high int
	size int
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

func push(q *queue, v any) {
	if q.high >= q.size-1 {
		fmt.Println("Очередь полна")
		return
	}
	q.high++
	q.s[q.high] = v
	if q.low == -1 {
		q.low = 0
	}
}

func pop(q *queue) any {
	if q.low == -1 || q.low > q.high {
		fmt.Println("Очередь пуста")
		return nil
	}
	val := q.s[q.low]
	q.low++
	return val
}

func main() {
	q := newQueue(5)
	push(q, 100)
	push(q, 200)
	push(q, 300)
	fmt.Println("Извлечено:", pop(q))
	fmt.Println("Извлечено:", pop(q))
	fmt.Println("Извлечено:", pop(q))
	fmt.Println("Извлечено:", pop(q))
}
