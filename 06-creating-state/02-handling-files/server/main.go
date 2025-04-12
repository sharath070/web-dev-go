package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request caught")
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("myfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.LimitReader()

		nf, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer nf.Close()

		// you're not getting a file, you're getting a stream of bytes,
		// and you decide what to do with it — save, rename, process, discard.
		size, err := io.Copy(nf, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("header: ", h)
		fmt.Println("size of file copied: ", size)
	}
}
