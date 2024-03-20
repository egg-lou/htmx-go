package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct{
	Title string
	Director string
}

func main() {
	fmt.Println("Server is running...")


	h1 := func (w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The GodFather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":4000", nil))
}