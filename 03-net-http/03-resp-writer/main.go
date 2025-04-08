package main

import (
	"fmt"
	"net/http"
)

type foo int

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("omai-wa-mo-shinde-iru", "Nani???")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>hello world</h1>")
}

func main() {
	var myfoo foo
	http.ListenAndServe(":8080", myfoo)
}
