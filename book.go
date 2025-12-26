package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("«%s», автор: %s (%d год)", b.Title, b.Author, b.Year)
}

func main() {
	
	book1 := Book{Title: "Война и мир", Author: "Лев Толстой", Year: 1869}
	book2 := Book{Title: "Преступление и наказание", Author: "Фёдор Достоевский", Year: 1866}

	fmt.Println(book1)
	fmt.Println(book2)
}
