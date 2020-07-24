package handler

import "github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/02-chain-of-responsibility/domain"

type Department interface {
	Execute(*domain.Patient)
	SetNext(Department)
}
