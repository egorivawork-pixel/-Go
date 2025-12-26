// Разработка многопоточного калькулятора

package main

import (
	"fmt"
	"time"
)

type CalcRequest struct {
	a, b     float64
	operator string
	resultCh chan float64
}

func calculator(requests chan CalcRequest) {
	for req := range requests {
		var result float64

		switch req.operator {
		case "+":
			result = req.a + req.b
		case "-":
			result = req.a - req.b
		case "*":
			result = req.a * req.b
		case "/":
			if req.b != 0 {
				result = req.a / req.b
			} else {
				fmt.Println("Ошибка: деление на ноль")
				req.resultCh <- 0
				continue
			}
		default:
			fmt.Println("Неизвестная операция:", req.operator)
			req.resultCh <- 0
			continue
		}

		req.resultCh <- result
	}
}

func client(a, b float64, op string, requests chan CalcRequest) {
	resultCh := make(chan float64)

	req := CalcRequest{a: a, b: b, operator: op, resultCh: resultCh}
	requests <- req

	result := <-resultCh
	fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n", a, op, b, result)
}

func main() {
	requests := make(chan CalcRequest)

	go calculator(requests)

	go client(10, 5, "+", requests)
	go client(10, 5, "-", requests)
	go client(10, 5, "*", requests)
	go client(10, 5, "/", requests)
	go client(10, 0, "/", requests)

	time.Sleep(2 * time.Second)
	fmt.Println("Все операции выполнены.")
}
