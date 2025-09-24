// Лабораторная работа №2
// Задание 2. Определение знака числа.
// Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero"


package main

import "fmt"

// Функция checkNumber принимает число и возвращает его характеристику
func checkNumber(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	}
	return "Zero"
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	// Вызов функции и вывод результата
	fmt.Println(checkNumber(num))
}
