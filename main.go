package main

import (
	"fmt"

	"github.com/marsonparulian/golang-lesson-struct-copied-in-methods/library"
)

func main() {
	lib := library.Library{}

	fmt.Println(lib)

	lib.AddABook(library.Book{Title: "Big Brown Bear"})

	fmt.Println(lib)
}
