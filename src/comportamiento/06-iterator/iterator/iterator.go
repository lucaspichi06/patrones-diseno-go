package iterator

import "github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/domain"

type Iterator interface {
	HasNext() bool
	GetNext() *domain.User
}
