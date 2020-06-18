package main

import (
	"fmt"
	"log"

	"github.com/lucaspichi06/patrones-diseno-go/src/estructurales/06-facade/facade"
)

func main() {
	showInit()

	to := "a@algo.com"
	comment := "muy buen video :)"
	token := "token-valido"
	user := "user-blog"

	f := facade.New(to, comment, token, user)
	if err := f.Comment(); err != nil {
		log.Fatal(err)
	}

	f.Notify()

	showFinish()
}

func showInit() {
	fmt.Println(`
	**************************
	* Bienvenido al programa *
	**************************
	`)
}

func showFinish() {
	fmt.Println(`
	**************************
	*  Gracias por utilizar  *
	**************************
	`)
}
