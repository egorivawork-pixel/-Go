// 1. Написание программы, которая будет выполнять три различные функции
// Каждая функция в отдельное горутине и с иметацией задержки при помощи time.Sleep()

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int, result chan int) {
	time.Sleep(2 * time.Second) 
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	result <- fact 
}

func randomNumbers(result chan int) {
	time.Sleep(1 * time.Second) 
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	result <- num
}


func sumSeries(n int, result chan int) {
	time.Sleep(3 * time.Second) 
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	result <- sum 
}

func main() {
	factorialResult := make(chan int)
	randomResult := make(chan int)
	sumResult := make(chan int)

	go factorial(5, factorialResult)
	go randomNumbers(randomResult)
	go sumSeries(10, sumResult)

	fact := <-factorialResult
	randomNum := <-randomResult
	sum := <-sumResult

	fmt.Printf("Факториал числа 5: %d\n", fact)
	fmt.Printf("Случайное число: %d\n", randomNum)
	fmt.Printf("Сумма числового ряда от 1 до 10: %d\n", sum)
}
