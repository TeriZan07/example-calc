package main

import (
	"net/http"

	"skeleton/src/app"
)

func main() {
	handlerCont := app.NewHandlers()
	mappings := app.BuildMappings(handlerCont)

	println("Listening to http://localhost:3000/ ")
	_ = http.ListenAndServe("127.0.0.1:3000", mappings)
}
