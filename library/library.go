package library

import "fmt"

type Book struct {
	Title string
}
type Library struct {
	LastAction string // Last action done
	books      []Book
}

func (lib Library) AddABook(book Book) {
	lib.LastAction = fmt.Sprintf("Book '%s' is added.", book.Title)
	lib.books = append(lib.books, book)
}
