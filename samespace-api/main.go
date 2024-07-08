package main

import (
	"log"
	"net/http"

	"samespace/db"
	"samespace/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.InitSession()
	defer db.Session.Close()

	r := mux.NewRouter()

	r.HandleFunc("/todos", handlers.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
