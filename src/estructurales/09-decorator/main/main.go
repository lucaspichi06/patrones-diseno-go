package main

import (
	"fmt"
	"os"

	decorator "github.com/lucaspichi06/patrones-diseno-go/src/estructurales/09-decorator"
)

func main() {
	route := decorator.NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scanf("%s", &path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	route.Exec(path)
}

func start(route *decorator.Route) {
	route.Add(decorator.NewLogRegistry(decorator.NewHandlerSayHello()), "/saludar") // haciendo esto, estoy decorando el NewHandlerSayHello()
	route.Add(decorator.NewLogRegistry(decorator.NewHandlerSayGoodbye()), "/despedirse")
}
