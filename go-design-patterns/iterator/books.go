package main

import "fmt"

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(Book) error) {
	var err error
	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

func (l *Library) createIterator() iterator {
	return NewBookIterator(l.Collection)
}

var lib *Library = &Library{
	Collection: []Book{
		{
			Name:      "War and Peace",
			Author:    "Leo Tolsotoy",
			PageCount: 864,
			Type:      HardCover,
		}, {
			Name:      "Crime and Punishment",
			Author:    "Leo Tolstoy",
			PageCount: 1225,
			Type:      PaperBack,
		},
	},
}
