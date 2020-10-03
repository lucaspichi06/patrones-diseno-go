package domain

type mediator interface {
	CanLand(Train) bool
	NotifyFree()
}
