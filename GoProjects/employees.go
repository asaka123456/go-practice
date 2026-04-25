package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   int
}

var commands = `
1 - Добавить сотрудника
2 - Удалить сотрудника
3 - Показать всех сотрудников
4 - Выйти
`

func main() {
	const size = 512
	empls := [size]*Employee{}
	empCount := 0

	for {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			if empCount >= size {
				fmt.Println("Достигнут максимум сотрудников (512)")
				continue
			}
			empl := new(Employee)
			fmt.Println("Имя:")
			fmt.Scanf("%s", &empl.Name)
			fmt.Println("Возраст:")
			fmt.Scanf("%d", &empl.Age)
			fmt.Println("Должность:")
			fmt.Scanf("%s", &empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scanf("%d", &empl.Salary)

			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					empCount++
					fmt.Println("Сотрудник добавлен!")
					break
				}
			}

		case 2:
			fmt.Println("Введите индекс сотрудника (0-511):")
			var idx int
			fmt.Scanf("%d", &idx)
			if idx >= 0 && idx < size && empls[idx] != nil {
				fmt.Println("Удаление:", empls[idx].Name)
				empls[idx] = nil
				empCount--
			} else {
				fmt.Println("Сотрудник не найден")
			}

		case 3:
			fmt.Println("\nСписок сотрудников")
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					e := empls[i]
					fmt.Printf("[%d] Имя: %s, Возраст: %d, Должность: %s, Зарплата: %d\n",
						i, e.Name, e.Age, e.Position, e.Salary)
				}
			}
			fmt.Println("----------------------------")

		case 4:
			fmt.Println("До свидания!")
			return
		}
	}
}
