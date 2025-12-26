// 2. Использование каналов для передачи данных:
// одна горутина реализирует числа фибоначи, другая их читает и выводит на экран.

package main

import (
	"fmt"
	"sync"
)


func fibonacci(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() 
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a 
		a, b = b, a+b
	}
	close(ch) 
}


func printFibonacci(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()       
	for num := range ch { 
		fmt.Println(num)
	}
}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go fibonacci(10, ch, &wg)

	wg.Add(1)
	go printFibonacci(ch, &wg)

	wg.Wait()
}
