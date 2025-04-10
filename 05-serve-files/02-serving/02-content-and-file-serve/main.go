package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog/pic", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}

func dog(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound) // 404
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound) // 404
		return
	}

	// this serves the file ans takes other info in `e` tag as cache
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
