package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

/*
	http.FileServer(http.Dir(".")) acts as a local sever or (live server in vscode)
	it servers anything and everything in the file including main.go
	This is the reason i couldn't access toby.jpg before directly in the
	img tag.... you need server those things seperately

	http.FileServer(http.Dir(".")) return an handler interface
	yeah the type that had ServeHTTP(ResponseWriter, *Request) method
*/
