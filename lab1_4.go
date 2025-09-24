// task4.go
// Задание 4: Арифметические операции с двумя целыми числами.

package main

import "fmt"

func main() {
	// Объявляем две переменные
	a, b := 15, 4

	// Выполняем и выводим основные операции
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)  // целочисленное деление
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // остаток от деления
}
