package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (h hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var h hotDog
	http.ListenAndServe(":8080", h)
}

/*
	ListenAndServe(addr string, handler Handler)
	this is a func that is used to listen to a port address and serve the contents
	in my example i made a type that inherits Handler interface
		type Handler interface {
			ServeHTTP(w http.ResponseWriter, r *http.Request)
		}
	since my type have this method it is also a handler
	this is why I can provide it as arg for `ListenAndServe`
*/
