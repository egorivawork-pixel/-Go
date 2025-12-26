package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch) // сигнал читателю: поток закончился
}

func printer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range ch { // range сам остановится, когда канал закроется
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go fibonacci(10, ch, &wg)

	wg.Add(1)
	go printer(ch, &wg)

	wg.Wait()
}