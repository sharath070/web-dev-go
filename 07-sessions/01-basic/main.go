package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.New()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			Secure:   true, // only valid for https
			HttpOnly: true, // cannot access cookie with javascript
		}
		http.SetCookie(w, c)
	}
	fmt.Fprintln(w, c)
}
