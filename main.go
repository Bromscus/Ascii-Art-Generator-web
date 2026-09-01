package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"web/ascii"
)

var (
	dataS  []byte
	dataSh []byte
	dataT  []byte
	result string
	data   []byte
)

type PageData struct {
	Result string
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return
	}

	pageData := PageData{
		Result: result,
	}

	tmpl.Execute(w, pageData)
}

func main() {
	var err error
	dataS, err = os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataT, err = os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataSh, err = os.ReadFile("shadow.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", asciiArt)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	choice := r.FormValue("styles")
	msg, isValid := ascii.IsValid(text)
	if !isValid {
		fmt.Fprintln(w, msg)
		return
	}

	switch choice {
	case "standard":
		data = dataS
	case "shadow":
		data = dataSh
	case "thinkertoy":
		data = dataT
	default:
		data = dataS
	}
	result := ascii.Process(text, data)
	fmt.Fprintln(w, result)
}
