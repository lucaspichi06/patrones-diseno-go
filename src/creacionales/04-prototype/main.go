package main

import (
	"fmt"
)

func main() {
	color := "Azul"
	phones := map[string]string{"home": "123456", "work": "789456"}
	p1 := prototype{"Lucas", 31, []string{"Ruso", "Fede"}, &color, phones}

	// copia del objeto original
	//p2 := p1
	p2 := p1.Clone()

	// cambio uno de los valores al objeto p2
	p2.age = 35
	p2.name = "Mariano"
	p2.friends[0] = "María" // al ser un slice, se modifica en ambos objetos. Debido a que un slice es un puntero a un array.
	p2.friends[1] = "Pedro"

	// Estos cambios sólo afectan al objeto p1 después de implementar el prototype
	color = "Rojo"
	phones["home"] = "1530222"

	// muestra la información
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
