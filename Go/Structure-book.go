package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Books
	book1.title = "Go Programming"
	book1.author = "Mahesh Kumar"
	book1.subject = "Go Programming Tutorial"
	book1.book_id = 567392

	book2 := Books{"Telecom Biling", "Zara Ali", "Telecom Billing Tutorial", 988492}

	printBook(book1)
	printBook(book2)
}
func printBook(book Books) {

	fmt.Printf("\nBook title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)

}

/*fmt.Printf("Book 1 title : %s\n", book1.title)
fmt.Printf("Book 1 author : %s\n", book1.author)
fmt.Printf("Book 1 subject : %s\n", book1.subject)
fmt.Printf("Book 1 book_id : %d\n", book1.book_id)

fmt.Printf("\nBook 2 title : %s\n", book2.title)
fmt.Printf("Book 2 author : %s\n", book2.author)
fmt.Printf("Book 2 subject : %s\n", book2.subject)
fmt.Printf("Book 2 book_id : %d\n", book2.book_id)*/
//}
