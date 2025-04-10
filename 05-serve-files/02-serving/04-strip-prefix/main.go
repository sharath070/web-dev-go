package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	/////////
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	/////////

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*
	http.StripPrefix()
	it strips the prefix of any url if it matches the pattern (here /resources)
	and redirects to a folder (here ./assets)

	here img src becomes "/resources/toby.jpg" to "./assets/toby.jpg"
*/

/*
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	this is the common pattern we see in code bases
	it's like you take the url that matches the pattern "/assets/"
	strip `assets` part and redirect to "./assets/toby.jpg" folder since it is being served
*/
