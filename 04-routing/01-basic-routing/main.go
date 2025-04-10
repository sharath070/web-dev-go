package main

import (
	"fmt"
	"net/http"
)

type foo int

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "doggy doggy doggy")
	case "/cat":
		fmt.Fprintln(w, "kitty kitty kitty")
	}
}

func main() {
	var f foo
	http.ListenAndServe(":8080", f)
}
