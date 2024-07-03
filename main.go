package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Buy Milk"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		temp1 := template.Must(template.ParseFiles("index.html"))

		temp1.Execute(w, data)
	}

	// fmt.Println("Helloo")
	http.HandleFunc("/", todosHandler)

	// log.Fatal here for if something goes wrong, it snaps out console logs the error.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
