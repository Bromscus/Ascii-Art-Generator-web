package main

import (
	"fmt"
	"log"
	"net/http"

	"web/functions"
)

func main() {
	http.HandleFunc("/", functions.Handler)
	http.HandleFunc("/ascii-art", functions.Handler)

	http.Handle(
		"/style/",
		http.StripPrefix("/style/", http.FileServer(http.Dir("style"))),
	)
	functions.LoadHashes()

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
