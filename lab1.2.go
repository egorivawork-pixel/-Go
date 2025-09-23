// task2.go
// Задание 2: Создать переменные различных типов и вывести их на экран.

package main

import "fmt"

func main() {
	// Объявление переменных с явным указанием типа
	var i int = 42         // целое число
	var f float64 = 3.14   // число с плавающей точкой
	var s string = "Hello" // строка
	var b bool = true      // логическая переменная

	// Выводим значения переменных
	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("string:", s)
	fmt.Println("bool:", b)
}
