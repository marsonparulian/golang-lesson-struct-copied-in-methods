package main

import (
	"fmt"
	"time"
)

type book struct {
	title string
}

type library struct {
	books []book
}

func (lib *library) addABook(book book) {
	fmt.Printf("Add book '%s'\n", book.title)
	lib.books = append(lib.books, book)
}

// Method below is mimicking time consuming operation.
// Method below is a Value receiver type.
func (lib *library) timeConsumingOperation() {
	// `Sleep` to mimick ` time-consuming operation.
	time.Sleep(30 * time.Millisecond)

	// Check the number of books
	fmt.Printf("Value receiver type. Found %d books.\n", len(lib.books))
}
func main() {
	lib := library{}

	// Goroutine mimicking a time consuming operation.
	go lib.timeConsumingOperation()

	// Add 2 books
	go lib.addABook(book{title: "Rain Rain Go Away"})
	go lib.addABook(book{title: "Rainbow After Rain"})

	// `Sleep` to make sure all goroutines are completed
	time.Sleep(2 * time.Second)

	fmt.Printf("Actual number of books : %d books.\n", len(lib.books))

}
