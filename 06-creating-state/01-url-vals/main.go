package main

import (
	"fmt"
	"net/http"
)

/*
	`2 types of forms`
	Form Method	Data Sent Where?	Typical Use Case					Secure for passwords?
	GET	In URL						Search, filters, pagination			❌ No, visible in URL
	POST							In request body	Login, signup,  	✅ Yes
									submit form

	`POST` form sends the data through `BODY`
	`GET` form sends the data through `URL`
*/

func main() {
	http.HandleFunc("/login", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		val := r.URL.Query().Get("dog") // get the val of dog from url
		fmt.Fprintf(w, "key = dog value = %s", val)

	case http.MethodPost:
		// expected to be used on post forms only
		email := r.FormValue("email") // if key email not present it returns empty string
		fmt.Fprintf(w, "your email: %s", email)
	}
}

// visit page => localhost:8080/login?email="babaBlackSheep@hotmail.com"

/*

	------ Use case ------			---- Need ParseForm()? ----
	Using req.FormValue(...)			❌ No
	Accessing req.Form["x"]				✅ Yes
	Using req.PostForm[...]				✅ Yes
	Parsing JSON body					❌ Use json.Decode

*/
