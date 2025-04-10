package main

import (
	"fmt"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "kitty kitty kitty")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
	// here nil represent the http is using default serve mux
}
