// Лабораторная работа №2
// Задание 4. Определение длины строки.
// Написать функцию, которая принимает строку и возвращает ее длину.


package main

import "fmt"

// Функция stringLength возвращает количество символов в строке
func stringLength(s string) int {
	return len(s)
}

func main() {
	var str string
	fmt.Print("Введите строку: ")
	// Используем Scanln, чтобы считать строку с пробелами
	fmt.Scanln(&str)

	// Выводим длину строки
	fmt.Println("Длина строки:", stringLength(str))
}
