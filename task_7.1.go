package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int, out chan<- int) {
	time.Sleep(2 * time.Second)
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	out <- f
}

func randomNumber(out chan<- int) {
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	out <- rand.Intn(100)
}

func sumSeries(n int, out chan<- int) {
	time.Sleep(3 * time.Second)
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	out <- s
}

func main() {
	factCh := make(chan int)
	randCh := make(chan int)
	sumCh := make(chan int)

	go factorial(5, factCh)
	go randomNumber(randCh)
	go sumSeries(10, sumCh)

	fact := <-factCh
	rn := <-randCh
	sum := <-sumCh

	fmt.Println("Факториал 5 =", fact)
	fmt.Println("Случайное число =", rn)
	fmt.Println("Сумма 1..10 =", sum)
}