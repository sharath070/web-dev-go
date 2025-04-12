package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = make(map[string]user)
var dbSession = make(map[string]string)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.New()
		c = &http.Cookie{Name: "session", Value: id.String()}
		http.SetCookie(w, c)
	}

	var u user
	if un, ok := dbSession[c.Value]; ok {
		u = dbUsers[un]
	}

	// process the submitted form
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSession[c.Value] = un
		dbUsers[un] = u
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// returning from this func is imp cuz even though we redirect the code continues
		// to execute till the end of this func
		return
	}

	un, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	fmt.Fprintln(w, u)
}
