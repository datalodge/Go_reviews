package main

import (
	"log"
	"net/http"

	"github.com/datalodge/go_reviews/cassandra"
	"github.com/datalodge/go_reviews/reviews"
	"github.com/gorilla/mux"
)

// main function

func main() {
	fs := http.FileServer(http.Dir("dist/"))
	CassandraConnection := cassandra.Session
	defer CassandraConnection.Close()
	http.Handle("/", http.StripPrefix("/", fs))
	router := mux.NewRouter()
	router.Handle("/{id}", fs)
	router.Handle("/", fs)
	router.Handle("/?", fs)
	router.HandleFunc("/api/reviews/{id}", reviews.Get).Methods("GET")
	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}
