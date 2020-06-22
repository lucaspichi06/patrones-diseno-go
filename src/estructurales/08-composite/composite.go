package main

// Item - esta interface deben implementarla todos aquellos objetos que van a formar nuestro árbol.-
type Item interface {
	Add(Item)
	String() string
	Price() int
}
