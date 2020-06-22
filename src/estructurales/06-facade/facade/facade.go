package facade

import (
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/06-facade/email"
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/06-facade/storage"
	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/06-facade/validate"
)

type Facade struct {
	// Facade debe conocer todos los subsistemas, y cada subsistema va a ser una propiedad de esta estructura.-
	to                  string
	comment             string
	validatorToken      validate.ValidatorToken
	validatorPermission validate.ValidatorPermission
	store               storage.Storage
	notificator         email.Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      validate.NewValidatorToken(token),
		validatorPermission: validate.NewValidatorPermission(user),
		store:               storage.NewStorage("mysql"),
		notificator:         email.NewEmail(),
	}
}

func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		return err
	}

	f.store.Save(f.comment)
	f.notificator.Send(f.to, f.comment)

	return nil
}

// puedo también crear funciones más pequeñas específicas para cierta función como por ejemplo Notify que únicamente envíe el mail.-
func (f *Facade) Notify() {
	f.notificator.Send(f.to, f.comment)
}
