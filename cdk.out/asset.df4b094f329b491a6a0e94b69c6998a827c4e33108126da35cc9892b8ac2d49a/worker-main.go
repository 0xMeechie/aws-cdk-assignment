package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Worker Server")
	http.HandleFunc("/", handleRoute)

	log.Fatalln(http.ListenAndServe(":80", nil))
}

func handleRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello from the Worker Task")
}
