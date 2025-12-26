// Запустить 1000 горутин на увелечение значения, с помощью мьютекса предотвратить гонку данных.
package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()

	counter++

	mu.Unlock()

	time.Sleep(time.Millisecond)
}

func main() {

	var wg sync.WaitGroup

	numGoroutines := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("Конечный счётчик с мьютексом:", counter)
}
