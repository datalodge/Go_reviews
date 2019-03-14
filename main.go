package main

import (
	"log"
	"net/http"

	"github.com/data_lodge/go_reviews/reviews"

	"github.com/data_lodge/go_reviews/cassandra"

	"github.com/gorilla/mux"
)

// main function

func main() {
	CassandraConnection := cassandra.Session
	defer CassandraConnection.Close()
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", http.StripPrefix("/", fs))
	router := mux.NewRouter()
	router.HandleFunc("/api/reviews/{id}", reviews.Get)
	router.Handle("/", fs)
	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}
