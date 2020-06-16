package connection

import "time"

// DBConnection es la intefaz para las conexiones de BD
type DBConnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	Close() error
}
