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
	if r.URL.Path != "/" {
		http.Error(w, "404: Status page not found Error", http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("styles")

	msg, valid := ascii.IsValid(text)
	if !valid {
		http.Error(w, msg, http.StatusBadRequest)
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
