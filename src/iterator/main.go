package main

type Book struct {
	name string
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	Next() *T
}

type BookShelf struct {
	books []*Book
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < len(bsi.bookShelf.books)
}

func (bsi *BookShelfIterator) Next() *Book {
	if !bsi.HasNext() {
		return nil
	}
	book := bsi.bookShelf.books[bsi.index]
	bsi.index++
	return book
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bookShelf, index: 0}
}

func (bs *BookShelf) Iterator() Iterator[Book] {
	return NewBookShelfIterator(bs)
}

func NewBookShelf() *BookShelf {
	bs := &BookShelf{}
	return bs
}

func main() {
	bookShelf := NewBookShelf()
	bookShelf.books = append(bookShelf.books, &Book{name: "Book 1"})
	bookShelf.books = append(bookShelf.books, &Book{name: "Book 2"})
	bookShelf.books = append(bookShelf.books, &Book{name: "Book 3"})

	var it Iterator[Book]
	it = bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		println(book.name)
	}
}
