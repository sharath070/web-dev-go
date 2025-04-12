package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	http.SetCookie(w, &http.Cookie{Name: "key", Value: "value"})
	fmt.Fprintln(w, "COOKIE SET")
	fmt.Fprintln(w, `<h1><a href="/read">Read Cookie</a></h1>`)
	fmt.Fprintln(w, `<h1><a href="/expire">Delete Cookie</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	c, err := r.Cookie("key")
	if err != nil {
		fmt.Fprintln(w, "NO COOKIE FOUND")
		fmt.Fprintln(w, `<h1><a href="/">Set Cookie</a></h1>`)
	} else {
		fmt.Fprintln(w, c)
		fmt.Fprintln(w, `<h1><a href="/expire">Delete Cookie</a></h1>`)
	}
}

func expire(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	c, err := r.Cookie("key")
	if err != nil {
		fmt.Fprintln(w, "NO COOKIE FOUND")
	} else {
		c.MaxAge = -1 // delete cookie
		http.SetCookie(w, c)
		fmt.Fprintln(w, "COOKIE DELETED")
	}

	fmt.Fprintln(w, `<h1><a href="/read">Read Cookie</a></h1>`)
	fmt.Fprintln(w, `<h1><a href="/">Set Cookie</a></h1>`)
}
