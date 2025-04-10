package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog/pic", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound) // 404
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}
