package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Task struct {
	line    string
	lineNum int
}

type Result struct {
	worker  int
	lineNum int
	text    string
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tasks {
		reversed := reverseString(t.line)
		results <- Result{
			worker:  id,
			lineNum: t.lineNum,
			text:    reversed,
		}
	}
}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&numWorkers)

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "
")

	tasks := make(chan Task)
	results := make(chan Result)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// закрываем results когда ВСЕ воркеры закончат
	go func() {
		wg.Wait()
		close(results)
	}()

	// отправляем задачи
	go func() {
		for i, line := range lines {
			line = strings.TrimRight(line, "
") // важно для Windows
			if strings.TrimSpace(line) == "" {
				continue
			}
			tasks <- Task{line: line, lineNum: i + 1}
		}
		close(tasks)
	}()

	var out strings.Builder
	for r := range results {
		out.WriteString(fmt.Sprintf("Worker %d (Line %d): %s
", r.worker, r.lineNum, r.text))
	}

	if err := os.WriteFile("output.txt", []byte(out.String()), 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Done. Saved to output.txt")
}