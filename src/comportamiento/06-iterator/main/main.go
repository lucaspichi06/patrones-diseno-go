package main

import (
	"fmt"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/domain"
	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/06-iterator/iterator"
)

func main() {
	user1 := &domain.User{
		Name: "a",
		Age:  30,
	}
	user2 := &domain.User{
		Name: "b",
		Age:  20,
	}
	userCollection := &iterator.UserCollection{
		Users: []*domain.User{user1, user2},
	}
	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
