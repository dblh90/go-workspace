package main

type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func NewBookIterator(books []Book) *BookIterator {
	return &BookIterator{
		books: books,
	}
}

func (bI *BookIterator) hasNext() bool {
	if bI.current < len(bI.books) {
		return true
	}
	return false
}

func (bI *BookIterator) next() *Book {
	if bI.hasNext() {
		bk := bI.books[bI.current]
		bI.current++
		return &bk
	}
	return nil
}
