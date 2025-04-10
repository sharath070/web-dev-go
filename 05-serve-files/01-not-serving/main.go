package main

import (
	"io"
	"net/http"
)

func cat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := `
	<!-- not serving from our server -->
	<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Himalayan_CAT.jpg/250px-Himalayan_CAT.jpg">
	`
	io.WriteString(w, content)
}

func main() {
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":8080", nil)
}
