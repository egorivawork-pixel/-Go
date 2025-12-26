package main

import "fmt"

func main() {

	var n int
	sum := 0

	fmt.Print("сколько чисел будете вводить? ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		sum += x
	}

	fmt.Println("сумма чисел =", sum)
}
