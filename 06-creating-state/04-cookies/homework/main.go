package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// using cookies track how many times the user has visited your site

func main() {
	http.HandleFunc("/", set)
	http.ListenAndServe(":8080", nil)
}

type state int

const (
	SET state = iota
	UPDATE
)

func set(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visited")
	if err == http.ErrNoCookie {
		c = &http.Cookie{Name: "visited", Value: "0"}
	}

	count, _ := strconv.Atoi(c.Value)
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)
	fmt.Fprintln(w, c)
}
