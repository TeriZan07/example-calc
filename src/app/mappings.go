package app

import (
	"net/http"
)

func BuildMappings(handlerContainer Handlers) *http.ServeMux {
	mappings := http.NewServeMux()

	//View mappings
	mappings.HandleFunc("/", handlerContainer.Index.Handle)
	mappings.HandleFunc("/error", handlerContainer.Error.Handle)

	//Api mappings
	mappings.HandleFunc("/api/sum", handlerContainer.Sum.Handle)
	mappings.HandleFunc("/api/mul", handlerContainer.Mul.Handle)

	//Serving scripts and styles
	mappings.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./src/front/css"))))
	mappings.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./src/front/js"))))

	return mappings
}
