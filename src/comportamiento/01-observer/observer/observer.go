package observer

// Observer interface del observer
type Observer interface {
	Notify(data string)
}
