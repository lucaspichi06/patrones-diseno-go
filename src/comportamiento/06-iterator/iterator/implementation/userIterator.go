package implementation

import "github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/domain"

type UserIterator struct {
	Index int
	Users []*domain.User
}

func (u *UserIterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false
}

func (u *UserIterator) GetNext() *domain.User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}
