package main

import "fmt"

type item struct {
	v    any
	next *item
}

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		fmt.Println("Индекс вне диапазона")
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

func remove(l *singlyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		fmt.Println("Индекс вне диапазона")
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		current := l.first
		for i := 0; i < idx-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		if idx == l.size-1 {
			l.last = current
		}
	}
	l.size--
}

func values(l *singlyLinkedList) []any {
	result := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

func main() {
	list := newSinglyLinkedList()
	add(list, "A")
	add(list, "B")
	add(list, "C")
	fmt.Println("Все значения:", values(list))
	fmt.Println("Получить индекс 1:", get(list, 1))
	remove(list, 0)
	fmt.Println("После удаления индекса 0:", values(list))
}
