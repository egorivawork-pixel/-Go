// 3. Применение select для управления каналами:
// • Создайте две горутины, одна из которых будет генерировать случайные числа, а другая
// — отправлять сообщения об их чётности/нечётности.
// • Используйте конструкцию select для приёма данных из обоих каналов и вывода
// результатов в консоль.
// • Продемонстрируйте, как select управляет многоканальными операциями.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для генерации случайных чисел
func generateRandomNumbers(ch chan int) {
	for {
		num := rand.Intn(100)   
		ch <- num               
		time.Sleep(time.Second) 
	}
}

func checkEvenOdd(ch chan int, result chan string) {
	for num := range ch {
		if num%2 == 0 {
			result <- fmt.Sprintf("%d — чётное", num)
		} else {
			result <- fmt.Sprintf("%d — нечётное", num)
		}
	}
}

func main() {
	randomNumbersCh := make(chan int)
	evenOddResultCh := make(chan string)

	go generateRandomNumbers(randomNumbersCh)

	go checkEvenOdd(randomNumbersCh, evenOddResultCh)

	for i := 0; i < 10; i++ { 
		select {
		case result := <-evenOddResultCh:
			fmt.Println(result)
		case num := <-randomNumbersCh:
			fmt.Println("Получено число:", num) 
		}
	}

	fmt.Println("Программа завершена.")
}
