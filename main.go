package main

import (
	"fmt"
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

	fmt.Println("Server running at http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
