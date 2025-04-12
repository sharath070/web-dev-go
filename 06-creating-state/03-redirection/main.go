package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/old-way", redirectOldWay)
	http.HandleFunc("/std-way", redirectStdWay)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func redirectStdWay(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func redirectOldWay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/std-way")
	w.WriteHeader(http.StatusSeeOther)
}

/*
	3 types of redirect
	http.StatusSeeOther 		 -> send `GET` request method to new url
	http.StatusTemporaryRedirect -> forward the same that you got request method to new url
	http.StatusMovedPermanently  -> resource permanently moved to new url
*/
