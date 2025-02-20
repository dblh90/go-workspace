package main

import "fmt"

func main() {
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
