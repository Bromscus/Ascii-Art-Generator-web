package functions

import (
	"html/template"
	"net/http"
	"os"

	"web/ascii"
)

type PageData struct {
	Result string
	Font   string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.Error(w, "404: Status page not found Error", http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404: Template Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/" && r.Method == http.MethodGet {
		err = templ.Execute(w, PageData{})
		if err != nil {
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	if r.URL.Path == "/ascii-art" && r.Method != http.MethodPost {
		http.Error(w, "405: Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("styles")

	msg, valid := ascii.IsValid(text)
	if !valid {
		data := PageData{
			Font:   font,
			Result: msg,
		}

		err = templ.Execute(w, data)
		if err != nil {
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	fontFile := ""

	switch font {
	case "standard":

		fontFile = "formats/standard.txt"

	case "shadow":

		fontFile = "formats/shadow.txt"

	case "thinkertoy":

		fontFile = "formats/thinkertoy.txt"

	default:
		http.Error(w, "400: Status Bad Request", http.StatusBadRequest)
		return
	}

	banner, err := os.ReadFile(fontFile)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404: Banner Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Font:   font,
		Result: ascii.Process(text, banner),
	}

	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}
}
