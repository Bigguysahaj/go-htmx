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
			Todo{Id: 1, Message: "Run Go server and HTMX client"},
			Todo{Id: 2, Message: "Video Editing"},
			Todo{Id: 3, Message: "Study theo next js"},
			Todo{Id: 4, Message: "Study theo 3D three js react stuff"},
			Todo{Id: 5, Message: "Study for iitm exam"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		temp1 := template.Must(template.ParseFiles("index.html"))

		temp1.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		// A variable message, that takes value from the form with label message
		message := r.PostFormValue("message")
		// Again parsing html in the new template
		temp1 := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)
		temp1.ExecuteTemplate(w, "todo-list-element", todo)
	}

	// fmt.Println("Helloo")
	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo/", addTodoHandler)

	// log.Fatal here for if something goes wrong, it snaps out console logs the error.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
