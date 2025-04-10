package main

import (
	"fmt"
	"net/http"
)

type foo int

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "doggy doggy doggy")
}

type bar int

func (b bar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "kitty kitty kitty")
}

func main() {
	var f foo
	var b bar

	mux := http.NewServeMux()

	mux.Handle("/dog/", f)
	// this will catch something that comes after /dog/... it can be /dog/somthing or /dog/something/else
	// it can also take /dog

	mux.Handle("/cat", b)
	// this will only take /cat... even if we provide /cat/ it gives back 404 err

	http.ListenAndServe(":8080", mux)
}
