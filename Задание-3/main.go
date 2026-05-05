package main

import (
	"fmt"

	"gitlab.com/mediasoft-internship/practice2024/less-3/linkedlist"
	"gitlab.com/mediasoft-internship/practice2024/less-3/matrix"
	"gitlab.com/mediasoft-internship/practice2024/less-3/queue"
	"gitlab.com/mediasoft-internship/practice2024/less-3/roman"
	"gitlab.com/mediasoft-internship/practice2024/less-3/stack"
)

func main() {
	// Стек
	s := stack.New[int](100)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Stack pop:", s.Pop())
	fmt.Println("Stack pop:", s.Pop())

	// Очередь
	q := queue.New[int](5)
	q.Push(100)
	q.Push(200)
	fmt.Println("Queue pop:", q.Pop())

	// Список
	l := linkedlist.New[string]()
	l.Add("A")
	l.Add("B")
	l.Add("C")
	fmt.Println("List values:", l.Values())

	// Римские цифры
	fmt.Println("Roman IX =", roman.ToArabic("IX"))

	// Матрица
	m := matrix.GenerateUnique(3, 3)
	fmt.Println("Matrix 3x3:")
	for _, row := range m {
		fmt.Println(row)
	}
}
