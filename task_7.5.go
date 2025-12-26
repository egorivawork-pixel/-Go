package main

import (
	"fmt"
	"sync"
)

type CalcRequest struct {
	a, b     float64
	op       string
	replyCh  chan float64
	errCh    chan error
}

func calculator(requests <-chan CalcRequest) {
	for r := range requests {
		var res float64
		switch r.op {
		case "+":
			res = r.a + r.b
		case "-":
			res = r.a - r.b
		case "*":
			res = r.a * r.b
		case "/":
			if r.b == 0 {
				r.errCh <- fmt.Errorf("деление на ноль")
				continue
			}
			res = r.a / r.b
		default:
			r.errCh <- fmt.Errorf("неизвестная операция: %s", r.op)
			continue
		}
		r.replyCh <- res
	}
}

func client(wg *sync.WaitGroup, requests chan<- CalcRequest, a, b float64, op string) {
	defer wg.Done()

	replyCh := make(chan float64, 1)
	errCh := make(chan error, 1)

	requests <- CalcRequest{a: a, b: b, op: op, replyCh: replyCh, errCh: errCh}

	select {
	case res := <-replyCh:
		fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, res)
	case err := <-errCh:
		fmt.Printf("%.2f %s %.2f -> ошибка: %v\n", a, op, b, err)
	}
}

func main() {
	requests := make(chan CalcRequest)
	go calculator(requests)

	var wg sync.WaitGroup
	wg.Add(5)

	go client(&wg, requests, 10, 5, "+")
	go client(&wg, requests, 10, 5, "-")
	go client(&wg, requests, 10, 5, "*")
	go client(&wg, requests, 10, 5, "/")
	go client(&wg, requests, 10, 0, "/")

	wg.Wait()
	close(requests)
	fmt.Println("Все операции выполнены.")
}