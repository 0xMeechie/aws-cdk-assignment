package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starting Service Task")
	http.HandleFunc("/", handleHome)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello From the service Task")
}
