package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./static")))
}

/*
	if there is a index.html present in the serving dir
	it will run by default on that route unless something is manually accessed
*/
