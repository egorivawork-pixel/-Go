// Лабораторная работа №2
// Задание 6. Среднее значение двух чисел.
// Написать функцию, которая принимает два целых числа и возвращает их среднее значение.


package main

import "fmt"

// Функция average возвращает среднее арифметическое двух чисел
func average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)

	fmt.Printf("Среднее значение = %.2f\n", average(a, b))
}
