package main

import (
	"fmt"
	"mathutils"
)

func main() {
	var число int

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&число)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	факториал := mathutils.Factorial(число)

	fmt.Printf("Факториал числа %d равен %d\n", число, факториал)
}
