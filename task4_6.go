package main

import "fmt"

func main() {
	var n int
	fmt.Print("введите количество элементов: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("массив в обратном порядке:")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
