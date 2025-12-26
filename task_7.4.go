package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup, useMutex bool) {
	defer wg.Done()

	if useMutex {
		mu.Lock()
		counter++
		mu.Unlock()
	} else {
		// специально без защиты — чтобы увидеть проблему
		counter++
	}
}

func main() {
	const goroutines = 1000

	// Поставь false, чтобы “выключить” мьютекс и посмотреть разницу
	useMutex := true

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go increment(&wg, useMutex)
	}
	wg.Wait()

	fmt.Println("useMutex =", useMutex, "counter =", counter)
}