package validate

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("El usuario no está autorizado para comentar")

type ValidatorPermission struct {
	userID string
}

func NewValidatorPermission(ID string) ValidatorPermission {
	return ValidatorPermission{userID: ID}
}

func (vp *ValidatorPermission) Validate() error {
	if vp.userID != "user-blog" {
		return ErrPermissionNotValid
	}

	fmt.Println("usuario autorizado")
	return nil

}
