// Лабораторная работа №2
// Задание 5. Структура Rectangle и метод вычисления площади.
// Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.



package main

import "fmt"

// Описание структуры прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод Area возвращает площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var w, h float64
	fmt.Print("Введите ширину и высоту прямоугольника: ")
	fmt.Scan(&w, &h)

	rect := Rectangle{Width: w, Height: h}
	fmt.Printf("Площадь прямоугольника = %.2f\n", rect.Area())
}
