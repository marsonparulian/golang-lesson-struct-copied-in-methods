package main

import (
	"fmt"
)

type book struct {
	title string
}

// Change the title of a book
func (b book) changeTitle(title string) {
	b.title = title
}

type library struct {
	books []book
}

func (lib library) addABook(book book) {
	lib.books = append(lib.books, book)
}

func main() {
	lib := library{}

	fmt.Println(lib)

	b := book{title: "Three little pigs"}
	b.changeTitle("Blue Birds")
	lib.addABook(b)

	fmt.Println(lib)
}
