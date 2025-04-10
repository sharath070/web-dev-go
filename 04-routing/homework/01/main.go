package main

import (
	"fmt"
	"net/http"
)

func defaultHF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bombardino Crocodilo...")
}

func catHF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Trippi Troppi... Troppa Trippa")
}

func meHF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sharath")
}

func main() {
	// here Handle takes route and Handler interface
	// we can make the method follow interface by putting it in http.HandlerFunc()
	http.Handle("/", http.HandlerFunc(defaultHF))
	http.Handle("/cat/", http.HandlerFunc(catHF))
	http.Handle("/me/", http.HandlerFunc(meHF))

	http.ListenAndServe(":8080", nil)
}
