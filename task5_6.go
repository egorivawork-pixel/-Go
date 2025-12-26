package main

import (
	"fmt"
)

// структура Book хранит информацию о книге
type Book struct {
	Title  string
	Author string
	Year   int
}

// реализуем интерфейс fmt.Stringer для структуры Book
func (b Book) String() string {
	return fmt.Sprintf("«%s», автор: %s (%d год)", b.Title, b.Author, b.Year)
}

func main() {
	// создаём несколько книг
	book1 := Book{Title: "Война и мир", Author: "Лев Толстой", Year: 1869}
	book2 := Book{Title: "Преступление и наказание", Author: "Фёдор Достоевский", Year: 1866}

	// выводим книги (автоматически вызовется метод String())
	fmt.Println(book1)
	fmt.Println(book2)
}
