package handlers

import "net/http"

type Handler interface {
	Handle(w http.ResponseWriter, request *http.Request)
}