package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

// Connect realiza la conexión a la BD Postgres.-
func (p *Postgres) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"postgres",
		"127.0.0.1",
		"5432",
		"postgres",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

// GetNow devuelve la fecha actual del servidor de base de datos.-
func (p *Postgres) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := p.db.QueryRow("select CURRENT_DATE").Scan(t)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: &v", err)
		return nil, err
	}
	return t, err
}

// Close cierra la conexión a la BD MySQL.-
func (p *Postgres) Close() error {
	return p.db.Close()
}
