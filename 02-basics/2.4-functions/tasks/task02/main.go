// Задание 2: Счётчик с начальным значением
//
// Напиши функцию makeCounter(start int) func() int, которая:
//   - принимает начальное значение
//   - возвращает функцию-счётчик
//   - каждый вызов возвращает следующее число, начиная с start
//
// Ожидаемый вывод:
//   Счётчик от 5:
//   5
//   6
//   7
//   Счётчик от 100:
//   100
//   101
//   Счётчики независимы - счётчик от 5 продолжает:
//   8
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeCounter(start int) func() int
// Подсказка: текущее значение храни в переменной внутри makeCounter,
// и обращайся к ней из возвращаемой функции (это и есть замыкание)

func main() {
	// TODO: создай два независимых счётчика и проверь их работу
	startCounter1 := 5
	fmt.Printf("Счётчик от %d:\n", startCounter1)
	counter1 := makeCounter(startCounter1)
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())

	startCounter2 := 100
	fmt.Printf("Счётчик от %d:\n", startCounter2)
	counter2 := makeCounter(startCounter2)
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

	fmt.Printf("Счётчик от %d:\n", startCounter1)
	fmt.Println(counter1())
	fmt.Println(counter1())
}

func makeCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}
