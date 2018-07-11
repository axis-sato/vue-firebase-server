package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
}


func private(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello private!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", private)

	log.Fatal(http.ListenAndServe(":8000", r))
}
