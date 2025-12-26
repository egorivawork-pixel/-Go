package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

type Task struct {
	line     string
	lineNum  int
	resultCh chan string
}

func worker(id int, tasks chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {

		reversed := reverseString(task.line)

		task.resultCh <- fmt.Sprintf("Worker %d (Line %d): %s", id, task.lineNum, reversed)
	}
}

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&numWorkers)

	inputFile := "input.txt"
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	tasks := make(chan Task, len(lines))

	results := make(chan string, len(lines))
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for i, line := range lines {
		line = strings.TrimRight(line, "\r")
		if line != "" {
			tasks <- Task{line: line, lineNum: i + 1, resultCh: results}
		}
	}

	close(tasks)

	wg.Wait()

	close(results)

	outputFile := "output.txt"
	outputData := ""

	for result := range results {
		outputData += result + "\n"
	}

	err = ioutil.WriteFile(outputFile, []byte(outputData), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Processing completed. Results saved in 'output.txt'.")
}
