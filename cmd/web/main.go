package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Print("Server started on port :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)

}
