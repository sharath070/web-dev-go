package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.txt"))
}

func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n")

	err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		fmt.Println(err)
	}
}
