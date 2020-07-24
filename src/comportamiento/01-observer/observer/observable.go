package observer

// Observable interface que tienen que implementar aquellos objetos que quieren ser observables.-
type Observable interface {
	AddObserver(name string, o Observer)
	RemoveObserver(name string)
	NotifyObservers()
}
