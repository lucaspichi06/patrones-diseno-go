package local

import "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/10-proxy/book"

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
