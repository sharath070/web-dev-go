package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n")

	tpl, err = tpl.ParseFiles("two.txt", "three.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		fmt.Println(err)
	}

	// only executes the first one it finds
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

}
