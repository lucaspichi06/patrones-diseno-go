package iterator

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/iterator/implementation"
)

type UserCollection struct {
	Users []*domain.User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &implementation.UserIterator{
		Users: u.Users,
	}
}
