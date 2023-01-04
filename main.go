package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! This app is running on Cloud Run!\n")
	})

	fmt.Printf("Server running on port 8080.\nYay! You're all set. Now, go to http://localhost:8080/ in your browser\nand you should see the message 'Hello, World!\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
