package routes

import (
	"noteservice/handlers"

	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/note", handlers.CreateNote).Methods("POST")
	r.HandleFunc("/note", handlers.ReadNotes).Methods("GET")

	return r
}
