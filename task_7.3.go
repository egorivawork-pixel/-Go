package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(numbersCh chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		numbersCh <- rand.Intn(100)
		time.Sleep(300 * time.Millisecond)
	}
	close(numbersCh)
}

func parity(numbersCh <-chan int, parityCh chan<- string) {
	for n := range numbersCh {
		if n%2 == 0 {
			parityCh <- fmt.Sprintf("%d — чётное", n)
		} else {
			parityCh <- fmt.Sprintf("%d — нечётное", n)
		}
	}
	close(parityCh)
}

func main() {
	numbersCh := make(chan int)
	parityCh := make(chan string)

	go gen(numbersCh)
	go parity(numbersCh, parityCh)

	// Пока хотя бы один канал не закрыт — будем слушать
	for numbersCh != nil || parityCh != nil {
		select {
		case n, ok := <-numbersCh:
			if !ok {
				numbersCh = nil // “отключаем” case, чтобы select его больше не выбирал
				continue
			}
			fmt.Println("Получено число:", n)

		case msg, ok := <-parityCh:
			if !ok {
				parityCh = nil
				continue
			}
			fmt.Println("Проверка:", msg)
		}
	}

	fmt.Println("Готово.")
}