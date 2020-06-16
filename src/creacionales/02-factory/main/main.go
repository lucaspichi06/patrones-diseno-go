package main

import (
	"fmt"
	"os"

	"github.com/lucaspichi06/patrones-diseno-go/src/creacionales/02-factory/factory"
)

func main() {
	var t int
	fmt.Print("Ingrese la conexión deseada:\n1-MySQL\n2-Postgres\n")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer la opción: %v", err)
		os.Exit(1) //detiene toda la ejecución del programa
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor no válido")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Printf("No se pudo crear la conexión: %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo consultar la fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))
	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexión: v%", err)
	}
}
