// task5.go
// Задание 5: Функция для вычисления суммы и разности двух чисел с плавающей запятой.

package main

import "fmt"

// sumAndDiff принимает два float64 и возвращает их сумму и разность
func sumAndDiff(a, b float64) (float64, float64) {
	return a + b, a - b
}

func main() {
	// Вызываем функцию с примерами
	sum, diff := sumAndDiff(10.5, 4.2)

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
}
