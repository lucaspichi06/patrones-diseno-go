package factory

import "github.com/lucaspichi06/patrones-diseno-go/src/creacionales/02-factory/connection"

// Factory devuelve un tipo DBConnection.-
func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.MySql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
