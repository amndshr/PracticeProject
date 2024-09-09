package main

import "fmt"

var books []book = []book{
	{
		title:     "Science101",
		author:    "Herr Fassbrau",
		ISBN:      "12345AB",
		available: true,
	},

	{
		title:     "Math101",
		author:    "Frau Fassbrau",
		ISBN:      "67890CD",
		available: false,
	},
}

func main() {
	libraryName := "Pulutlib"
	display(libraryName)

	// create a newBook variable

	newBook := book{
		title:     "Golang101",
		author:    "Davud Safarov",
		ISBN:      "09876EF",
		available: true,
	}

	addNewBook(newBook)

	displayAllBook()

	addBookFromUser()

	displayAllBook()
}

type book struct {
	title     string
	author    string
	ISBN      string
	available bool
}

func display(libraryName string) {
	fmt.Println("Welcome to", libraryName)
}

func addNewBook(newBook book) {
	books = append(books, newBook)
}

func displayAllBook() {
	for i := 0; i < len(books); i++ {
		b := books[i]
		fmt.Print(i+1, ".", b.title)

		if b.available {
			fmt.Println(" - available")
		} else {
			fmt.Println(" - not available")
		}

	}
}

func addBookFromUser() {
	userBook := book{}

	fmt.Println("Insert title of the book: ")
	fmt.Scanln(&userBook.title)

	fmt.Println("Insert author of the book: ")
	fmt.Scanln(&userBook.author)

	fmt.Println("Insert book's ISBN: ")
	fmt.Scanln(&userBook.ISBN)

	fmt.Println("Is the book ready to order: true/false")
	fmt.Scanln(&userBook.available)

	addNewBook(userBook)
}

// {Science101 Herr Fassbrau 12345AB true}
// {Math101 Frau Fassbrau 67890CD false}
//

// Science101
// Math101
