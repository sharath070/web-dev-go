package main

import (
	"fmt"
	"log"
	"net/http"
)

// Request gives all the info of the client req
// (yeah the one we saw in tcp server from before when accessed from web)

func main() {

}

// This is used for the forms that is submitted using post req (usually in templates)
type PostFormFoo int

func (f PostFormFoo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // first this that should be done for processing forms
	if err != nil {
		log.Fatalln(err)
	}

	formVal := r.Form // type: map[string]([]string) => map of str to slice of str
	fmt.Println(formVal)
}
