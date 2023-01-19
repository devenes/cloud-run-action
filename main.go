package main

import (
	"fmt"           // package for input and output operations
	"html/template" // package for template parsing
	"net/http"      // package for HTTP server
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// parse the template file
		tmpl, err := template.ParseFiles("template.html")

		if err != nil {
			// if there is an error, send an error message to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// execute the template and write the result to the response writer
		err = tmpl.Execute(w, nil)

		if err != nil {
			// if there is an error, send an error message to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on http://localhost:8080") // print a message to the console
	http.ListenAndServe(":8080", nil)                 // start the server and listen on port 8080
}
