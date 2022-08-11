# Avoid Using Value Receiver Type in Golang Structs Methods

## About this article
This article assumes the readers already know basic knowledge of `struct`, `methods`, and `goroutine` in Go programming language.

The main purpose of this article is as reminder for myself about the use of value/pointer receiver type in struct methods (Go programming language).
I decided to write about this, even though this is a very simple thing, 
after I spend more then 1 hour to debug an issue caused by the use of value receiver type in struct methods.

In an example below I will show the issue caused by the use of valuke receiver but without showing any warnings.

## Receiver Types In Struct Methods

In Golang we define methods for a struct similar to methods in object-oriented programming language.
The difference with OO language is in Golang there are 2 types of of receiver : `value` and `pointer` receiver types.

Pointer receiver type will copy the pointer of the 'object/struct' in memory. This is the same with methods in OO programming language where any changes made in the methods will change the actual 'object/struct' in memory. 
Example of pointer receiver type is as below :

```Go
// Define a structtype Book struct {
	Title string
}

    // A pointer receiver type method. 
	// Marked with the asterisk (*) sign.
func(b *Book) ChangeTitle(title string) {
  b.Title = title
}
```

 Value receiver type will copy the 'object/struct' data. Any changes made to the object/struct' in methods will not change the actual 'object/struct' in memory.
Example of value receiver type is below : 

```Go
// Value receiver type. 
// Marked with no asterisk (*) sign, in the struct receiver. 
func (b Book) ChangeTitle(title string) {
    b.Title = title
}
```

In Golang we only use the value receiver if the method do not need to make any changes to the struct data.
However that is not true for all cases. In this article I will show an example where we need to use pointer receiver in a 'read-only'method.

## Value Pointer Type

### Example of `ineffective assignment` Warning
```Go
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
```

The Go language server (vscode extension), `gopls`, will show `ineffective-assignment` warning to the assignment in `changeTitle` and `addABook` methods.
This is very helpful to prevent any bugs in case we forgot to use pointer receiver type.

### Example in `goroutine` implementation

There is also other condition where pointer receiver is needed, but value receiver is used, yet no warnings are shown.
This situation arise with the use of `goroutine`s, like in the code below :

```
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
func (lib library) timeConsumingOperation() {
	// `Sleep` to mimick ` time-consuming operation.
	time.Sleep(5 * time.Millisecond)

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
```

Running code above will produce a false result.
The `Library` in the `timeConsumingOperation` will have the outdated version 
(since the methods were run as `goroutine`).

```
Add book 'Rainbow After Rain'
Add book 'Rain Rain Go Away'
Value receiver type. Found 0 books.
Actual number of books : 2 books.
```

## Pointer Receiver Type
The use of pointer receiver type in struct methods is mandatory if we want data integrity. 
Of course there are smaller number of cases where we need to use value receiver type.
Modification of earlier example, with pointer receiver type on the `timeConsumingOperation`, can be seen below (added an asterisk(*)).
```
func (lib *library) timeConsumingOperation() {
```
Using pointer type in `timeConsumingOperation` will
points the `lib` to the same data in the memory, assuring data integrity.
Running the modified code will produce : 

```Go
Add book 'Rainbow After Rain'
Add book 'Rain Rain Go Away'
Value receiver type. Found 2 books.
Actual number of books : 2 books.
```

## Conclusion
Using pointer receiver in methods, instead of value receiver types, will assure data integrity.
I suggest that we should , maybe except in few cases, to always use pointer receiver types to avoid potential bugs
that are not catched by tools like the case above.

Here is a link of [Code xample repository in Github.com][code-repo].

Hope this article can be useful and thanks for reading


[code-repo]: https://github.com/marsonparulian/golang-lesson-struct-copied-in-methods


